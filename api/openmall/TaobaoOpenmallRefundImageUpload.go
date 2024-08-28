package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundImageUpload OpenMall退款图片上传
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
func TaobaoOpenmallRefundImageUpload(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundImageUploadAPIRequest, resp *openmall.TaobaoOpenmallRefundImageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
