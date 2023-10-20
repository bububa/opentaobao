package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpeffectproductsingleget 单个产品的报表
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
func Alibabascbpeffectproductsingleget(clt *core.SDKClient, req *scbp.AlibabascbpeffectproductsinglegetAPIRequest, session string) (*scbp.AlibabascbpeffectproductsinglegetAPIResponse, error) {
	var resp scbp.AlibabascbpeffectproductsinglegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
