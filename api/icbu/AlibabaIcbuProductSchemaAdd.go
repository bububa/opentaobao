package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductschemaadd （新）商品发布新接口
// alibaba.icbu.product.schema.add
//
// 提供发布ICBU商品的入口
func Alibabaicbuproductschemaadd(clt *core.SDKClient, req *icbu.AlibabaicbuproductschemaaddAPIRequest, session string) (*icbu.AlibabaicbuproductschemaaddAPIResponse, error) {
	var resp icbu.AlibabaicbuproductschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
