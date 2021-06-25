package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
gifting消息获取 APIRequest
taobao.istore.gifting.message.get

该api通过参数查询对应的gifting消息
*/
type TaobaoIstoreGiftingMessageGetRequest struct {
    model.Params

    // 消息查询条件
    giftMessageBizCondition   *GiftMessageBizCondition 

}

func NewTaobaoIstoreGiftingMessageGetRequest() *TaobaoIstoreGiftingMessageGetRequest{
    return &TaobaoIstoreGiftingMessageGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIstoreGiftingMessageGetRequest) GetApiMethodName() string {
    return "taobao.istore.gifting.message.get"
}

func (r TaobaoIstoreGiftingMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIstoreGiftingMessageGetRequest) SetGiftMessageBizCondition(giftMessageBizCondition *GiftMessageBizCondition) error {
    r.giftMessageBizCondition = giftMessageBizCondition
    r.Set("gift_message_biz_condition", giftMessageBizCondition)
    return nil
}

func (r TaobaoIstoreGiftingMessageGetRequest) GetGiftMessageBizCondition() *GiftMessageBizCondition {
    return r.giftMessageBizCondition
}

