package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotRechargeCardAPIRequest 按终端号订购增值业务 API请求
// alibaba.aliqin.fc.iot.rechargeCard
//
// 按终端号订购增值业务
type AlibabaAliqinFcIotRechargeCardAPIRequest struct {
	model.Params
	// 外部id,用来做幂等
	_outRechargeId string
	// 外部计费号类型：写‘ICCID’
	_billSource string
	// iccid的值
	_billReal string
	// ICCID
	_iccid string
	// 流量包offerId
	_offerId string
	// 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
	_effCode string
	// yyyy-MM-dd HH:mm:ss
	_effTime string
}

// NewAlibabaAliqinFcIotRechargeCardRequest 初始化AlibabaAliqinFcIotRechargeCardAPIRequest对象
func NewAlibabaAliqinFcIotRechargeCardRequest() *AlibabaAliqinFcIotRechargeCardAPIRequest {
	return &AlibabaAliqinFcIotRechargeCardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.rechargeCard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRechargeId is OutRechargeId Setter
// 外部id,用来做幂等
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetBillSource is BillSource Setter
// 外部计费号类型：写‘ICCID’
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetBillSource(_billSource string) error {
	r._billSource = _billSource
	r.Set("bill_source", _billSource)
	return nil
}

// GetBillSource BillSource Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetBillSource() string {
	return r._billSource
}

// SetBillReal is BillReal Setter
// iccid的值
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetBillReal(_billReal string) error {
	r._billReal = _billReal
	r.Set("bill_real", _billReal)
	return nil
}

// GetBillReal BillReal Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetBillReal() string {
	return r._billReal
}

// SetIccid is Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetIccid() string {
	return r._iccid
}

// SetOfferId is OfferId Setter
// 流量包offerId
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetOfferId(_offerId string) error {
	r._offerId = _offerId
	r.Set("offer_id", _offerId)
	return nil
}

// GetOfferId OfferId Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetOfferId() string {
	return r._offerId
}

// SetEffCode is EffCode Setter
// 生效时间,1000,立即生效; AUTO_ORD,下周期自动续订
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetEffCode(_effCode string) error {
	r._effCode = _effCode
	r.Set("eff_code", _effCode)
	return nil
}

// GetEffCode EffCode Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetEffCode() string {
	return r._effCode
}

// SetEffTime is EffTime Setter
// yyyy-MM-dd HH:mm:ss
func (r *AlibabaAliqinFcIotRechargeCardAPIRequest) SetEffTime(_effTime string) error {
	r._effTime = _effTime
	r.Set("eff_time", _effTime)
	return nil
}

// GetEffTime EffTime Getter
func (r AlibabaAliqinFcIotRechargeCardAPIRequest) GetEffTime() string {
	return r._effTime
}
