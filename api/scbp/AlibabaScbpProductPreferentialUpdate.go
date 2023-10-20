package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpproductpreferentialupdate 设置P4P产品优先推广状态
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
func Alibabascbpproductpreferentialupdate(clt *core.SDKClient, req *scbp.AlibabascbpproductpreferentialupdateAPIRequest, session string) (*scbp.AlibabascbpproductpreferentialupdateAPIResponse, error) {
	var resp scbp.AlibabascbpproductpreferentialupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
