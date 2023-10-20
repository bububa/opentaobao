package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagpuschemaupdate 产品更新接口
// alibaba.gpu.schema.update
//
// 产品更新接口
func Alibabagpuschemaupdate(clt *core.SDKClient, req *product.AlibabagpuschemaupdateAPIRequest, session string) (*product.AlibabagpuschemaupdateAPIResponse, error) {
	var resp product.AlibabagpuschemaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
