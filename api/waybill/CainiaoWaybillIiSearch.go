package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaowaybilliisearch 查询面单服务订购及面单使用情况
// cainiao.waybill.ii.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
func Cainiaowaybilliisearch(clt *core.SDKClient, req *waybill.CainiaowaybilliisearchAPIRequest, session string) (*waybill.CainiaowaybilliisearchAPIResponse, error) {
	var resp waybill.CainiaowaybilliisearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
