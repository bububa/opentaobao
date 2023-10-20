package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundimageuploadAPIRequest OpenMall退款图片上传 API请求
// taobao.openmall.refund.image.upload
//
// OpenMall退款图片上传
type TaobaoopenmallrefundimageuploadAPIRequest struct {
	model.Params
	// 渠道商Nick
	_distributor string
	// 上传图片，必须为jpg或png格式，建议小于2M
	_image *model.File
	// 该图片归属的退款单ID
	_refundId int64
}

// NewTaobaoopenmallrefundimageuploadRequest 初始化TaobaoopenmallrefundimageuploadAPIRequest对象
func NewTaobaoopenmallrefundimageuploadRequest() *TaobaoopenmallrefundimageuploadAPIRequest {
	return &TaobaoopenmallrefundimageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.image.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道商Nick
func (r *TaobaoopenmallrefundimageuploadAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetImage is Image Setter
// 上传图片，必须为jpg或png格式，建议小于2M
func (r *TaobaoopenmallrefundimageuploadAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetImage() *model.File {
	return r._image
}

// SetRefundId is RefundId Setter
// 该图片归属的退款单ID
func (r *TaobaoopenmallrefundimageuploadAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundimageuploadAPIRequest) GetRefundId() int64 {
	return r._refundId
}
