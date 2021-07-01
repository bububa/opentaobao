package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenMall退款图片上传 API请求
taobao.openmall.refund.image.upload

OpenMall退款图片上传
*/
type TaobaoOpenmallRefundImageUploadAPIRequest struct {
    model.Params
    // 上传图片，必须为jpg或png格式，建议小于2M
    _image   *model.File
    // 渠道商Nick
    _distributor   string
    // 该图片归属的退款单ID
    _refundId   int64
}

// 初始化TaobaoOpenmallRefundImageUploadAPIRequest对象
func NewTaobaoOpenmallRefundImageUploadRequest() *TaobaoOpenmallRefundImageUploadAPIRequest{
    return &TaobaoOpenmallRefundImageUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Image Setter
// 上传图片，必须为jpg或png格式，建议小于2M
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetImage(_image *model.File) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetImage() *model.File {
    return r._image
}
// Distributor Setter
// 渠道商Nick
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetDistributor() string {
    return r._distributor
}
// RefundId Setter
// 该图片归属的退款单ID
func (r *TaobaoOpenmallRefundImageUploadAPIRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundImageUploadAPIRequest) GetRefundId() int64 {
    return r._refundId
}
