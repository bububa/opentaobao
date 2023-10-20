package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// Alibabaicbuproductschemaupdate （新）商品发布增量更新接口
// alibaba.icbu.product.schema.update
//
// 商品更新接口，方式为增量更新，只更新传入的字段
func Alibabaicbuproductschemaupdate(clt *core.SDKClient, req *icbu.AlibabaicbuproductschemaupdateAPIRequest, session string) (*icbu.AlibabaicbuproductschemaupdateAPIResponse, error) {
	var resp icbu.AlibabaicbuproductschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
