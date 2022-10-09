// Twitter 的 Snowflake 算法的实现
/*
Snowflake 的结构如下（每部分用-分开）：
0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000

- 最高位是符号位，正数是0，负数是1，id一般是正数，因此最高位固定是0;
- 41位时间戳(毫秒级)，注意：41位时间戳不是存储当时时间的时间戳，而是存储时间戳的差值(当前时间戳-开始时间戳)，开始时间一般指定为项目启动时间，由程序指定，可以使用69年：(1<<41)/(1000*60*60*24*365)
- 10位的机器相关位，可以部署在1024个节点，包括5位的datacenterID(数据中心id)和5位的workerID(工作机器id)
- 12位系列号，毫秒内的计数。12位的计数顺序号支持每个节点每毫秒（同一机器，同一时间戳）产生4096个ID序号

*/

package muuid

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	sequenceMask   = 1<<12 - 1
	workerMask     = 1<<5 - 1
	dataCenterMask = 1<<5 - 1

	workerLeftShift     = 12
	dataCenterLeftShift = 17
	timestampLeftShift  = 22
)

// SnowFlake 雪花算法
type SnowFlake struct {
	mutex sync.Mutex

	sequence     int16
	dataCenterID uint8
	workerID     uint8

	lastTimeStamp int64 // 上次生成ID的时间戳(毫秒)

	startTime time.Time
}

// NewWith 给定开始时间和可选的 dataCenterID 和 workerID（注意两者的顺序）
// 如果ids没传，则使用machineID
func NewWith(startTime time.Time, ids ...uint8) *SnowFlake {
	var dataCenterID, workerID uint8

	idLen := len(ids)
	switch {
	case idLen >= 2:
		dataCenterID, workerID = ids[0], ids[1]
	case idLen == 1:
		dataCenterID, workerID = ids[0], ids[0]
	default:
		dataCenterID, workerID = machineID()
	}

	return &SnowFlake{
		startTime:    startTime.UTC(),
		dataCenterID: dataCenterID & dataCenterMask,
		workerID:     workerID & workerMask,
	}
}

// New 创建新的SnowFlake
func New() *SnowFlake {
	startTime := time.Date(2022, 9, 29, 0, 0, 0, 0, time.UTC)

	return NewWith(startTime)
}

// NextID 获取下一个ID
func (s *SnowFlake) NextID() int64 {
	now := time.Now().UTC()
	millisecond := now.UnixMilli()
	if millisecond < s.lastTimeStamp {
		panic("Clock moved backwards, Refusing to generate id")
	}

	s.mutex.Lock()
	// 同一毫秒，进行毫秒内序号递增
	if millisecond == s.lastTimeStamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		// 当前毫秒内序号用完，堵塞到下一毫秒
		if s.sequence == 0 {
			for millisecond <= s.lastTimeStamp {
				millisecond = getMillisecond()
			}
		}
	} else {
		// 时间戳改变，毫秒内序号重置
		s.sequence = 0
	}
	s.lastTimeStamp = millisecond
	sequence := s.sequence
	s.mutex.Unlock()

	elaspedMillisecond := millisecond - s.startTime.UnixMilli()

	return elaspedMillisecond<<timestampLeftShift |
		int64(s.dataCenterID)<<dataCenterMask |
		int64(s.workerID)<<workerLeftShift |
		int64(sequence)
}

// String 覆写string方法
func (s *SnowFlake) String() string {
	return fmt.Sprintf("start_time:%s, data_center_id:%d, worker_id:%d, sequence:%d",
		s.startTime, s.dataCenterID, s.workerID, s.sequence)
}

func machineID() (uint8, uint8) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return 0, 0
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if ip != nil {
			return ip[2], ip[3]
		}
	}

	return 0, 0
}

// getMillisecond 获取当前UTC时间的时间戳（毫秒表示）
func getMillisecond() int64 {
	return time.Now().UTC().UnixMilli()
}
