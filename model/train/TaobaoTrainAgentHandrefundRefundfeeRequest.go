package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商手动退款接口 APIRequest
taobao.train.agent.handrefund.refundfee

火车票代理商手动退款接口
*/
type TaobaoTrainAgentHandrefundRefundfeeRequest struct {
    model.Params

    // 主订单id
    mainBizOrderId   int64 

    // 外部订单号
    outTradeNo   string 

    // 退款金额,单位为分
    refundFee   int64 

}

func NewTaobaoTrainAgentHandrefundRefundfeeRequest() *TaobaoTrainAgentHandrefundRefundfeeRequest{
    return &TaobaoTrainAgentHandrefundRefundfeeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetApiMethodName() string {
    return "taobao.train.agent.handrefund.refundfee"
}

func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetMainBizOrderId(mainBizOrderId int64) error {
    r.mainBizOrderId = mainBizOrderId
    r.Set("main_biz_order_id", mainBizOrderId)
    return nil
}

func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetMainBizOrderId() int64 {
    return r.mainBizOrderId
}

func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetRefundFee() int64 {
    return r.refundFee
}

