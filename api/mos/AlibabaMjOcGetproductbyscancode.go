package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocgetproductbyscancode POS商品查询接口
// alibaba.mj.oc.getproductbyscancode
//
// 此API用于在银泰商场中，POS端扫码获取商品信息
func Alibabamjocgetproductbyscancode(clt *core.SDKClient, req *mos.AlibabamjocgetproductbyscancodeAPIRequest, session string) (*mos.AlibabamjocgetproductbyscancodeAPIResponse, error) {
	var resp mos.AlibabamjocgetproductbyscancodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
