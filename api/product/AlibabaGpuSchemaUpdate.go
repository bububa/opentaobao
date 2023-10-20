package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGpuSchemaUpdate 产品更新接口
// alibaba.gpu.schema.update
//
// 产品更新接口
func AlibabaGpuSchemaUpdate(clt *core.SDKClient, req *product.AlibabaGpuSchemaUpdateAPIRequest, resp *product.AlibabaGpuSchemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
