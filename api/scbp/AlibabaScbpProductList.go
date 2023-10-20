package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpproductlist 查询P4P产品
// alibaba.scbp.product.list
//
// 查询P4P产品
func Alibabascbpproductlist(clt *core.SDKClient, req *scbp.AlibabascbpproductlistAPIRequest, session string) (*scbp.AlibabascbpproductlistAPIResponse, error) {
	var resp scbp.AlibabascbpproductlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
