package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductschemaget （新）ICBU商品发布schema接口
// alibaba.icbu.product.schema.get
//
// 获取ICBU商品发布的页面规则和填写字段，适用于新发商品
func Alibabaicbuproductschemaget(clt *core.SDKClient, req *icbu.AlibabaicbuproductschemagetAPIRequest, session string) (*icbu.AlibabaicbuproductschemagetAPIResponse, error) {
	var resp icbu.AlibabaicbuproductschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
