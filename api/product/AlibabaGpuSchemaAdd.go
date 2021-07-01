package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* AlibabaGpuSchemaAdd
使用schema文件发布产品
alibaba.gpu.schema.add

使用Schema文件发布一个产品 */
func AlibabaGpuSchemaAdd(clt *core.SDKClient, req *product.AlibabaGpuSchemaAddAPIRequest, session string) (*product.AlibabaGpuSchemaAddAPIResponse, error) {
	var resp product.AlibabaGpuSchemaAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
