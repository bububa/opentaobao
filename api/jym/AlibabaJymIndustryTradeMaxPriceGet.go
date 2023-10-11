package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryTradeMaxPriceGet 获取交易猫单个游戏渠道帐号交易成功最高价
// alibaba.jym.industry.trade.max.price.get
//
// 获取交易猫单个游戏渠道帐号交易成功最高价
func AlibabaJymIndustryTradeMaxPriceGet(clt *core.SDKClient, req *jym.AlibabaJymIndustryTradeMaxPriceGetAPIRequest, session string) (*jym.AlibabaJymIndustryTradeMaxPriceGetAPIResponse, error) {
	var resp jym.AlibabaJymIndustryTradeMaxPriceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
