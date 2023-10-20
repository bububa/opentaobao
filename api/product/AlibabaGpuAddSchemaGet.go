package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagpuaddschemaget 获取产品发布规则接口
// alibaba.gpu.add.schema.get
//
// 获取产品发布规则接口
func Alibabagpuaddschemaget(clt *core.SDKClient, req *product.AlibabagpuaddschemagetAPIRequest, session string) (*product.AlibabagpuaddschemagetAPIResponse, error) {
	var resp product.AlibabagpuaddschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
