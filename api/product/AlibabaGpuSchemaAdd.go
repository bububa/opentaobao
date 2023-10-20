package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagpuschemaadd 使用schema文件发布产品
// alibaba.gpu.schema.add
//
// 使用Schema文件发布一个产品
func Alibabagpuschemaadd(clt *core.SDKClient, req *product.AlibabagpuschemaaddAPIRequest, session string) (*product.AlibabagpuschemaaddAPIResponse, error) {
	var resp product.AlibabagpuschemaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
