package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductschemarenderdraft （新）渲染草稿商品数据
// alibaba.icbu.product.schema.render.draft
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
func Alibabaicbuproductschemarenderdraft(clt *core.SDKClient, req *icbu.AlibabaicbuproductschemarenderdraftAPIRequest, session string) (*icbu.AlibabaicbuproductschemarenderdraftAPIResponse, error) {
	var resp icbu.AlibabaicbuproductschemarenderdraftAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
