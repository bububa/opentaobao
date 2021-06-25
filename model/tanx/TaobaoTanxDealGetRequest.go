package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对外部dsp提供交易id查询接口 APIRequest
taobao.tanx.deal.get

对外部dsp提供交易id查询接口
*/
type TaobaoTanxDealGetRequest struct {
    model.Params

    // dsp用户id
    dspId   int64 

    // 交易id
    dealId   int64 

    // 1970年到现在的时间，毫秒
    signTime   int64 

    // 验证token
    token   string 

}

func NewTaobaoTanxDealGetRequest() *TaobaoTanxDealGetRequest{
    return &TaobaoTanxDealGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxDealGetRequest) GetApiMethodName() string {
    return "taobao.tanx.deal.get"
}

func (r TaobaoTanxDealGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxDealGetRequest) SetDspId(dspId int64) error {
    r.dspId = dspId
    r.Set("dsp_id", dspId)
    return nil
}

func (r TaobaoTanxDealGetRequest) GetDspId() int64 {
    return r.dspId
}

func (r *TaobaoTanxDealGetRequest) SetDealId(dealId int64) error {
    r.dealId = dealId
    r.Set("deal_id", dealId)
    return nil
}

func (r TaobaoTanxDealGetRequest) GetDealId() int64 {
    return r.dealId
}

func (r *TaobaoTanxDealGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxDealGetRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxDealGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxDealGetRequest) GetToken() string {
    return r.token
}

