package film

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退票接口 APIRequest
taobao.film.data.third.party.refund.order

淘票票第三方退票接口
*/
type TaobaoFilmDataThirdPartyRefundOrderRequest struct {
    model.Params

    // 淘宝账号ID，此ID是一串数字。可自行百度查看如何获取或者咨询淘票票技术人员提供
    userId   int64 

    // 淘票票分配的渠道码
    platform   int64 

    // 退票身份ID，用于标识一个购票用户的身份，该参数需要跟锁座接口的ext_order_id参数一致，否则下单会失败。外部渠道需保证该参数的唯一性及准确性，下单出票接口会利用该参数做冥等性判断，如果由于外部渠道自身传入的参数有问题而导致的下单出票接口返回的结果有误，需要外部渠道自己承担损失
    extUserId   string 

    // 退款时候需要传入第三方的订单号。外部渠道需保证该参数的唯一性和准确性
    extOrderId   string 

    // 下单时返回的淘宝订单号参数
    tbOrderId   int64 

    // 退款金额，以分为单位，为指定的退款订单的金额
    refundAmount   int64 

    // 退款服务费，目前都为0
    refundServiceFee   int64 

    // 目前可以暂时不填参数
    params   string 

}

func NewTaobaoFilmDataThirdPartyRefundOrderRequest() *TaobaoFilmDataThirdPartyRefundOrderRequest{
    return &TaobaoFilmDataThirdPartyRefundOrderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetApiMethodName() string {
    return "taobao.film.data.third.party.refund.order"
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetUserId() int64 {
    return r.userId
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetPlatform() int64 {
    return r.platform
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetExtUserId(extUserId string) error {
    r.extUserId = extUserId
    r.Set("ext_user_id", extUserId)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetExtUserId() string {
    return r.extUserId
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetExtOrderId(extOrderId string) error {
    r.extOrderId = extOrderId
    r.Set("ext_order_id", extOrderId)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetExtOrderId() string {
    return r.extOrderId
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetTbOrderId(tbOrderId int64) error {
    r.tbOrderId = tbOrderId
    r.Set("tb_order_id", tbOrderId)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetTbOrderId() int64 {
    return r.tbOrderId
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetRefundAmount(refundAmount int64) error {
    r.refundAmount = refundAmount
    r.Set("refund_amount", refundAmount)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetRefundAmount() int64 {
    return r.refundAmount
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetRefundServiceFee(refundServiceFee int64) error {
    r.refundServiceFee = refundServiceFee
    r.Set("refund_service_fee", refundServiceFee)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetRefundServiceFee() int64 {
    return r.refundServiceFee
}

func (r *TaobaoFilmDataThirdPartyRefundOrderRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r TaobaoFilmDataThirdPartyRefundOrderRequest) GetParams() string {
    return r.params
}

