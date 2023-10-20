package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryTradeMaxPriceGet 获取交易猫单个游戏渠道帐号交易成功最高价
// alibaba.jym.industry.trade.max.price.get
//
// 获取交易猫单个游戏渠道帐号交易成功最高价
func AlibabaJymIndustryTradeMaxPriceGet(clt *core.SDKClient, req *jym.AlibabaJymIndustryTradeMaxPriceGetAPIRequest, resp *jym.AlibabaJymIndustryTradeMaxPriceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
