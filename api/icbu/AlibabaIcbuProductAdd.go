package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductadd 发布产品
// alibaba.icbu.product.add
//
// 发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品
func Alibabaicbuproductadd(clt *core.SDKClient, req *icbu.AlibabaicbuproductaddAPIRequest, session string) (*icbu.AlibabaicbuproductaddAPIResponse, error) {
	var resp icbu.AlibabaicbuproductaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
