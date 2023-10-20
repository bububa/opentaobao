package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustrytrademaxpriceget 获取交易猫单个游戏渠道帐号交易成功最高价
// alibaba.jym.industry.trade.max.price.get
//
// 获取交易猫单个游戏渠道帐号交易成功最高价
func Alibabajymindustrytrademaxpriceget(clt *core.SDKClient, req *jym.AlibabajymindustrytrademaxpricegetAPIRequest, session string) (*jym.AlibabajymindustrytrademaxpricegetAPIResponse, error) {
	var resp jym.AlibabajymindustrytrademaxpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
