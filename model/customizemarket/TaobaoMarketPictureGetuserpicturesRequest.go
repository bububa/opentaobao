package customizemarket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
读取用户上传图片 API请求
taobao.market.picture.getuserpictures

商家通过用户信息，获取用户上传的
*/
type TaobaoMarketPictureGetuserpicturesRequest struct {
    model.Params
    // 订单ID
    _orderId   int64
}

// 初始化TaobaoMarketPictureGetuserpicturesRequest对象
func NewTaobaoMarketPictureGetuserpicturesRequest() *TaobaoMarketPictureGetuserpicturesRequest{
    return &TaobaoMarketPictureGetuserpicturesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMarketPictureGetuserpicturesRequest) GetApiMethodName() string {
    return "taobao.market.picture.getuserpictures"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMarketPictureGetuserpicturesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *TaobaoMarketPictureGetuserpicturesRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoMarketPictureGetuserpicturesRequest) GetOrderId() int64 {
    return r._orderId
}
