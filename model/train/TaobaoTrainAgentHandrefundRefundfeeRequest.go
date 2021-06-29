package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商手动退款接口 API请求
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

// 初始化TaobaoTrainAgentHandrefundRefundfeeRequest对象
func NewTaobaoTrainAgentHandrefundRefundfeeRequest() *TaobaoTrainAgentHandrefundRefundfeeRequest{
    return &TaobaoTrainAgentHandrefundRefundfeeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetApiMethodName() string {
    return "taobao.train.agent.handrefund.refundfee"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainBizOrderId Setter
// 主订单id
func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetMainBizOrderId(mainBizOrderId int64) error {
    r.mainBizOrderId = mainBizOrderId
    r.Set("main_biz_order_id", mainBizOrderId)
    return nil
}

// MainBizOrderId Getter
func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetMainBizOrderId() int64 {
    return r.mainBizOrderId
}
// OutTradeNo Setter
// 外部订单号
func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetOutTradeNo() string {
    return r.outTradeNo
}
// RefundFee Setter
// 退款金额,单位为分
func (r *TaobaoTrainAgentHandrefundRefundfeeRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoTrainAgentHandrefundRefundfeeRequest) GetRefundFee() int64 {
    return r.refundFee
}
