## thirdweb-api

## doc:
1. https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json

### API拆分
1. 部署分账合约: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deploySplit
2. 部署合约
  + 721: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deployNFTDrop
  + 1155: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deployEditionDrop
3. 设置版税: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Contract-Royalties/operation/setDefaultRoyaltyInfo
4. 设置掉落条件: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC721/operation/setClaimConditions

### cv使用到的接口
1. 部署collection专辑合约
  + 创建thirdweb合约
    - https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deployNFTDrop
    - https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deployEditionDrop
  + 创建分账合约: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Deploy/operation/deploySplit
  + 设置版税: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/Contract-Royalties/operation/setDefaultRoyaltyInfo
2. 揭示延迟铸造的721的nft: 
3. 设置nft购买条件，此处可能已经是复合调用
  + 721: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC721/operation/setClaimConditions
  + 1155: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC1155/operation/setClaimConditions
4. 批量铸造
  + 721: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC721/operation/mintBatchTo
  + 1155: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC1155/operation/mintBatchTo
5. 延迟铸造
  + 721: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC721/operation/lazyMint
  + 1155: https://redocly.github.io/redoc/?url=https://tw-engine.culturevault.com/json#tag/ERC1155/operation/lazyMint