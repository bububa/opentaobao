package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按终端号订购增值业务 APIRequest
alibaba.aliqin.fc.iot.rechargeCard

按终端号订购增值业务
*/
type AlibabaAliqinFcIotRechargeCardRequest struct {
    model.Params

    // 外部计费号类型：写‘ICCID’
    billSource   string 

    // iccid的值
    billReal   string 

    // 流量包offerId
    offerId   string 

    // 外部id,用来做幂等
    outRechargeId   string 

    // ICCID
    iccid   string 

    // 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
    effCode   string 

    // yyyy-MM-dd HH:mm:ss
    effTime   string 

}

func NewAlibabaAliqinFcIotRechargeCardRequest() *AlibabaAliqinFcIotRechargeCardRequest{
    return &AlibabaAliqinFcIotRechargeCardRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.rechargeCard"
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillSource() string {
    return r.billSource
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillReal() string {
    return r.billReal
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOfferId(offerId string) error {
    r.offerId = offerId
    r.Set("offer_id", offerId)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetOfferId() string {
    return r.offerId
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetOutRechargeId() string {
    return r.outRechargeId
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetIccid() string {
    return r.iccid
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffCode(effCode string) error {
    r.effCode = effCode
    r.Set("eff_code", effCode)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffCode() string {
    return r.effCode
}

func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffTime(effTime string) error {
    r.effTime = effTime
    r.Set("eff_time", effTime)
    return nil
}

func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffTime() string {
    return r.effTime
}

