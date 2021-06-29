package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gifting消息获取 API请求
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
type TaobaoIstoreGiftingMessageGetRequest struct {
    model.Params
    // 消息查询条件
    _giftMessageBizCondition   *GiftMessageBizCondition
}

// 初始化TaobaoIstoreGiftingMessageGetRequest对象
func NewTaobaoIstoreGiftingMessageGetRequest() *TaobaoIstoreGiftingMessageGetRequest{
    return &TaobaoIstoreGiftingMessageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIstoreGiftingMessageGetRequest) GetApiMethodName() string {
    return "taobao.istore.gifting.message.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIstoreGiftingMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GiftMessageBizCondition Setter
// 消息查询条件
func (r *TaobaoIstoreGiftingMessageGetRequest) SetGiftMessageBizCondition(_giftMessageBizCondition *GiftMessageBizCondition) error {
    r._giftMessageBizCondition = _giftMessageBizCondition
    r.Set("gift_message_biz_condition", _giftMessageBizCondition)
    return nil
}

// GiftMessageBizCondition Getter
func (r TaobaoIstoreGiftingMessageGetRequest) GetGiftMessageBizCondition() *GiftMessageBizCondition {
    return r._giftMessageBizCondition
}
