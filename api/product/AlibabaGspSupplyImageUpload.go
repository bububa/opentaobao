package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagspsupplyimageupload gsp图片上传
// alibaba.gsp.supply.image.upload
//
// 上传图片至目标海外平台的素材空间
func Alibabagspsupplyimageupload(clt *core.SDKClient, req *product.AlibabagspsupplyimageuploadAPIRequest, session string) (*product.AlibabagspsupplyimageuploadAPIResponse, error) {
	var resp product.AlibabagspsupplyimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
