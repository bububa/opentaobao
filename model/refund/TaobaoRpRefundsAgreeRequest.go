package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同意退款 APIRequest
taobao.rp.refunds.agree

卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
*/
type TaobaoRpRefundsAgreeRequest struct {
    model.Params

    // 短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。
    code   string 

    // 退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。
    refundInfos   string 

}

func NewTaobaoRpRefundsAgreeRequest() *TaobaoRpRefundsAgreeRequest{
    return &TaobaoRpRefundsAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRpRefundsAgreeRequest) GetApiMethodName() string {
    return "taobao.rp.refunds.agree"
}

func (r TaobaoRpRefundsAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRpRefundsAgreeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoRpRefundsAgreeRequest) GetCode() string {
    return r.code
}

func (r *TaobaoRpRefundsAgreeRequest) SetRefundInfos(refundInfos string) error {
    r.refundInfos = refundInfos
    r.Set("refund_infos", refundInfos)
    return nil
}

func (r TaobaoRpRefundsAgreeRequest) GetRefundInfos() string {
    return r.refundInfos
}

