package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaRender （新）获取商品信息
// alibaba.icbu.product.schema.render
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
func AlibabaIcbuProductSchemaRender(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaRenderAPIRequest, session string) (*icbu.AlibabaIcbuProductSchemaRenderAPIResponse, error) {
	var resp icbu.AlibabaIcbuProductSchemaRenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
