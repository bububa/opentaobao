package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGpuSchemaUpdate 产品更新接口
// alibaba.gpu.schema.update
//
// 产品更新接口
func AlibabaGpuSchemaUpdate(clt *core.SDKClient, req *product.AlibabaGpuSchemaUpdateAPIRequest, session string) (*product.AlibabaGpuSchemaUpdateAPIResponse, error) {
	var resp product.AlibabaGpuSchemaUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
