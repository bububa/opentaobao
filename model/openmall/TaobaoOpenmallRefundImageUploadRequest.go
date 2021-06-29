package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenMall退款图片上传 APIRequest
taobao.openmall.refund.image.upload

OpenMall退款图片上传
*/
type TaobaoOpenmallRefundImageUploadRequest struct {
    model.Params

    // 上传图片，必须为jpg或png格式，建议小于2M
    image   []*model.File 

    // 渠道商Nick
    distributor   string 

    // 该图片归属的退款单ID
    refundId   int64 

}

func NewTaobaoOpenmallRefundImageUploadRequest() *TaobaoOpenmallRefundImageUploadRequest{
    return &TaobaoOpenmallRefundImageUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallRefundImageUploadRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.image.upload"
}

func (r TaobaoOpenmallRefundImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallRefundImageUploadRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoOpenmallRefundImageUploadRequest) GetImage() []*model.File {
    return r.image
}

func (r *TaobaoOpenmallRefundImageUploadRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

func (r TaobaoOpenmallRefundImageUploadRequest) GetDistributor() string {
    return r.distributor
}

func (r *TaobaoOpenmallRefundImageUploadRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoOpenmallRefundImageUploadRequest) GetRefundId() int64 {
    return r.refundId
}

