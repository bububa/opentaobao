package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaGspSupplyImageUpload gsp图片上传
// alibaba.gsp.supply.image.upload
//
// 上传图片至目标海外平台的素材空间
func AlibabaGspSupplyImageUpload(clt *core.SDKClient, req *product.AlibabaGspSupplyImageUploadAPIRequest, resp *product.AlibabaGspSupplyImageUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
