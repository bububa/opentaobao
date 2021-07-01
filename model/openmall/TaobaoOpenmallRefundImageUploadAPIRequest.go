package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundImageUploadAPIRequest
OpenMall退款图片上传 API请求
taobao.openmall.refund.image.upload

OpenMall退款图片上传 */
type TaobaoOpenmallRefundImageUploadAPIRequest struct {
	model.Params
	// 上传图片，必须为jpg或png格式，建议小于2M
	_image *model.File
	// 渠道商Nick
	_distributor string
	// 该图片归属的退款单ID
	_refundId int64
}

// New
