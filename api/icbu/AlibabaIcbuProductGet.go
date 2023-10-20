package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductget 获得单个商品详情
// alibaba.icbu.product.get
//
// 获取商品详情
func Alibabaicbuproductget(clt *core.SDKClient, req *icbu.AlibabaicbuproductgetAPIRequest, session string) (*icbu.AlibabaicbuproductgetAPIResponse, error) {
	var resp icbu.AlibabaicbuproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
