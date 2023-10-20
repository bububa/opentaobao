package admarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/admarket"
)

// Yunosadmarketmaterialaudit 广告平台创意审核
// yunos.admarket.material.audit
//
// 用于厂商上报广告平台审核结果
func Yunosadmarketmaterialaudit(clt *core.SDKClient, req *admarket.YunosadmarketmaterialauditAPIRequest, session string) (*admarket.YunosadmarketmaterialauditAPIResponse, error) {
	var resp admarket.YunosadmarketmaterialauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
