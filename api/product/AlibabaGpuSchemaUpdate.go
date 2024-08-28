package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGpuSchemaUpdate 产品更新接口
// alibaba.gpu.schema.update
//
// 产品更新接口
func AlibabaGpuSchemaUpdate(ctx context.Context, clt *core.SDKClient, req *product.AlibabaGpuSchemaUpdateAPIRequest, resp *product.AlibabaGpuSchemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
