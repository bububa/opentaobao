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
    _billSource   string
    // iccid的值
    _billReal   string
    // 流量包offerId
    _offerId   string
    // 外部id,用来做幂等
    _outRechargeId   string
    // ICCID
    _iccid   string
    // 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
    _effCode   string
    // yyyy-MM-dd HH:mm:ss
    _effTime   string
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
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillSource(_billSource string) error {
    r._billSource = _billSource
    r.Set("bill_source", _billSource)
    return nil
}

// BillSource Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillSource() string {
    return r._billSource
}
// BillReal Setter
// iccid的值
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetBillReal(_billReal string) error {
    r._billReal = _billReal
    r.Set("bill_real", _billReal)
    return nil
}

// BillReal Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetBillReal() string {
    return r._billReal
}
// OfferId Setter
// 流量包offerId
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOfferId(_offerId string) error {
    r._offerId = _offerId
    r.Set("offer_id", _offerId)
    return nil
}

// OfferId Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetOfferId() string {
    return r._offerId
}
// OutRechargeId Setter
// 外部id,用来做幂等
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetOutRechargeId(_outRechargeId string) error {
    r._outRechargeId = _outRechargeId
    r.Set("out_recharge_id", _outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetOutRechargeId() string {
    return r._outRechargeId
}
// Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetIccid() string {
    return r._iccid
}
// EffCode Setter
// 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffCode(_effCode string) error {
    r._effCode = _effCode
    r.Set("eff_code", _effCode)
    return nil
}

// EffCode Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffCode() string {
    return r._effCode
}
// EffTime Setter
// yyyy-MM-dd HH:mm:ss
func (r *AlibabaAliqinFcIotRechargeCardRequest) SetEffTime(_effTime string) error {
    r._effTime = _effTime
    r.Set("eff_time", _effTime)
    return nil
}

// EffTime Getter
func (r AlibabaAliqinFcIotRechargeCardRequest) GetEffTime() string {
    return r._effTime
}
