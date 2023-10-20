package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstoreitemfromonlineitem 基于新模型商品id查询摊位子品id
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
func Tmallnrtstoreitemfromonlineitem(clt *core.SDKClient, req *nrt.TmallnrtstoreitemfromonlineitemAPIRequest, session string) (*nrt.TmallnrtstoreitemfromonlineitemAPIResponse, error) {
	var resp nrt.TmallnrtstoreitemfromonlineitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
