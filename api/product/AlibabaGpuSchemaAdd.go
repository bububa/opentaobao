package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGpuSchemaAdd 使用schema文件发布产品
// alibaba.gpu.schema.add
//
// 使用Schema文件发布一个产品
func AlibabaGpuSchemaAdd(ctx context.Context, clt *core.SDKClient, req *product.AlibabaGpuSchemaAddAPIRequest, resp *product.AlibabaGpuSchemaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
