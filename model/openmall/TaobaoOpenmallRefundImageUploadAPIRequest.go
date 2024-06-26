package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundImageUploadAPIRequest OpenMall退款图片上传 API请求
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
type TaobaoOpenmallRefundImageUploadAPIRequest struct {
	model.Params
	// 渠道商Nick
	_distributor string
	// 上传图片，必须为jpg或png格式，建议小于2M
	_image *model.File
	// 该图片归属的退款单ID
	_refundId int64
}

// NewTaobaoOpenmallRefundImageUploadRequest 初始化TaobaoOpenmallRefundImageUploadAPIRequest对象
func NewTaobaoOpenmallRefundImageUploadRequest() *TaobaoOpenmallRefundImageUploadAPIRequest {
	return &TaobaoOpenmallRefundImageUploadAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) Reset() {
	r._distributor = ""
	r._image = nil
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道商Nick
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetImage is Image Setter
// 上传图片，必须为jpg或png格式，建议小于2M
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetRefundId is RefundId Setter
// 该图片归属的退款单ID
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoOpenmallRefundImageUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallRefundImageUploadRequest()
	},
}

// GetTaobaoOpenmallRefundImageUploadRequest 从 sync.Pool 获取 TaobaoOpenmallRefundImageUploadAPIRequest
func GetTaobaoOpenmallRefundImageUploadAPIRequest() *TaobaoOpenmallRefundImageUploadAPIRequest {
	return poolTaobaoOpenmallRefundImageUploadAPIRequest.Get().(*TaobaoOpenmallRefundImageUploadAPIRequest)
}

// ReleaseTaobaoOpenmallRefundImageUploadAPIRequest 将 TaobaoOpenmallRefundImageUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallRefundImageUploadAPIRequest(v *TaobaoOpenmallRefundImageUploadAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallRefundImageUploadAPIRequest.Put(v)
}
