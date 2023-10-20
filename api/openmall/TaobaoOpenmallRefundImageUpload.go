package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundImageUpload OpenMall退款图片上传
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
func TaobaoOpenmallRefundImageUpload(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundImageUploadAPIRequest, resp *openmall.TaobaoOpenmallRefundImageUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
