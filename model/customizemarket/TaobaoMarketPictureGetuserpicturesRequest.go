package customizemarket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
读取用户上传图片 APIRequest
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的
*/
type TaobaoMarketPictureGetuserpicturesRequest struct {
    model.Params

    // 订单ID
    orderId   int64 

}

func NewTaobaoMarketPictureGetuserpicturesRequest() *TaobaoMarketPictureGetuserpicturesRequest{
    return &TaobaoMarketPictureGetuserpicturesRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMarketPictureGetuserpicturesRequest) GetApiMethodName() string {
    return "taobao.market.picture.getuserpictures"
}

func (r TaobaoMarketPictureGetuserpicturesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMarketPictureGetuserpicturesRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoMarketPictureGetuserpicturesRequest) GetOrderId() int64 {
    return r.orderId
}

