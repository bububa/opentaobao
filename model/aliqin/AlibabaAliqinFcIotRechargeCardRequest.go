package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按终端号订购增值业务 API请求
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

// 初始化AlibabaAliqinFcIotRechargeCardRequest对象
func NewAlibabaAliqinFcIotRechargeCardRequest() *AlibabaAliqinFcIotRechargeCardRequest{
    return &AlibabaAliqinFcIotRechargeCardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotRechargeCardRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.rechargeCard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotRechargeCardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillSource Setter
// 外部计费号类型：写‘ICCID’
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillSource(billSource string) error {
    r.billSource = billSource
    r.Set("bill_source", billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillSource() string {
    return r.billSource
}
// BillReal Setter
// iccid的值
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillReal(billReal string) error {
    r.billReal = billReal
    r.Set("bill_real", billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillReal() string {
    return r.billReal
}
// OfferId Setter
// 流量包offerId
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOfferId(offerId string) error {
    r.offerId = offerId
    r.Set("offer_id", offerId)
    return nil
}

// OfferId Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetOfferId() string {
    return r.offerId
}
// OutRechargeId Setter
// 外部id,用来做幂等
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetOutRechargeId() string {
    return r.outRechargeId
}
// Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetIccid() string {
    return r.iccid
}
// EffCode Setter
// 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffCode(effCode string) error {
    r.effCode = effCode
    r.Set("eff_code", effCode)
    return nil
}

// EffCode Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffCode() string {
    return r.effCode
}
// EffTime Setter
// yyyy-MM-dd HH:mm:ss
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffTime(effTime string) error {
    r.effTime = effTime
    r.Set("eff_time", effTime)
    return nil
}

// EffTime Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffTime() string {
    return r.effTime
}
