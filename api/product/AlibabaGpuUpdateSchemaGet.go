package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagpuupdateschemaget 获取产品编辑schema规则的接口
// alibaba.gpu.update.schema.get
//
// 获取产品编辑schema规则的接口
func Alibabagpuupdateschemaget(clt *core.SDKClient, req *product.AlibabagpuupdateschemagetAPIRequest, session string) (*product.AlibabagpuupdateschemagetAPIResponse, error) {
	var resp product.AlibabagpuupdateschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
