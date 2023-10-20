package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductschemaadddraft （新）ICBU商品发布草稿
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
func Alibabaicbuproductschemaadddraft(clt *core.SDKClient, req *icbu.AlibabaicbuproductschemaadddraftAPIRequest, session string) (*icbu.AlibabaicbuproductschemaadddraftAPIResponse, error) {
	var resp icbu.AlibabaicbuproductschemaadddraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
