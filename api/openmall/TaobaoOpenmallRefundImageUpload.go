package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundImageUpload OpenMall退款图片上传
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
func TaobaoOpenmallRefundImageUpload(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundImageUploadAPIRequest, session string) (*openmall.TaobaoOpenmallRefundImageUploadAPIResponse, error) {
	var resp openmall.TaobaoOpenmallRefundImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
