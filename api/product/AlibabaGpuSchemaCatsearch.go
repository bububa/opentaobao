package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGpuSchemaCatsearch 按类目查询spu接口
// alibaba.gpu.schema.catsearch
//
// 按类目查询spu的schema接口
func AlibabaGpuSchemaCatsearch(clt *core.SDKClient, req *product.AlibabaGpuSchemaCatsearchAPIRequest, resp *product.AlibabaGpuSchemaCatsearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
