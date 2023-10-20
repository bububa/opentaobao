package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallrefundimageupload OpenMall退款图片上传
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
func Taobaoopenmallrefundimageupload(clt *core.SDKClient, req *openmall.TaobaoopenmallrefundimageuploadAPIRequest, session string) (*openmall.TaobaoopenmallrefundimageuploadAPIResponse, error) {
	var resp openmall.TaobaoopenmallrefundimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
