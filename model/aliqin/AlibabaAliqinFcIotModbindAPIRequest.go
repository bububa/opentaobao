package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotModbindAPIRequest 物联网绑定/换绑API API请求
// alibaba.aliqin.fc.iot.modbind
//
// 支持用户的设备的换绑和解绑操作
type AlibabaAliqinFcIotModbindAPIRequest struct {
	model.Params
	// chgBind：换绑；unBind：解绑
	_opionType string
	// 物联卡和iccid保持一致
	_billReal string
	// iccid （20位）
	_iccid string
	// 换绑的时候必传，换的新设备imei
	_newimei string
	// 目前绑定的设备imei
	_imei string
	// 物联卡业务：若无特殊为ICCID
	_billSource string
	// 若无特殊物联卡传入122
	_midPatChannel string
}

// NewAlibabaAliqinFcIotModbindRequest 初始化AlibabaAliqinFcIotModbindAPIRequest对象
func NewAlibabaAliqinFcIotModbindRequest() *AlibabaAliqinFcIotModbindAPIRequest {
	return &AlibabaAliqinFcIotModbindAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcIotModbindAPIRequest) Reset() {
	r._opionType = ""
	r._billReal = ""
	r._iccid = ""
	r._newimei = ""
	r._imei = ""
	r._billSource = ""
	r._midPatChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotModbindAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.modbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotModbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotModbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpionType is OpionType Setter
// chgBind：换绑；unBind：解绑
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetOpionType(_opionType string) error {
	r._opionType = _opionType
	r.Set("opion_type", _opionType)
	return nil
}

// GetOpionType OpionType Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetOpionType() string {
	return r._opionType
}

// SetBillReal is BillReal Setter
// 物联卡和iccid保持一致
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetBillReal(_billReal string) error {
	r._billReal = _billReal
	r.Set("bill_real", _billReal)
	return nil
}

// GetBillReal BillReal Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetBillReal() string {
	return r._billReal
}

// SetIccid is Iccid Setter
// iccid （20位）
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetIccid() string {
	return r._iccid
}

// SetNewimei is Newimei Setter
// 换绑的时候必传，换的新设备imei
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetNewimei(_newimei string) error {
	r._newimei = _newimei
	r.Set("newimei", _newimei)
	return nil
}

// GetNewimei Newimei Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetNewimei() string {
	return r._newimei
}

// SetImei is Imei Setter
// 目前绑定的设备imei
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetImei() string {
	return r._imei
}

// SetBillSource is BillSource Setter
// 物联卡业务：若无特殊为ICCID
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetBillSource(_billSource string) error {
	r._billSource = _billSource
	r.Set("bill_source", _billSource)
	return nil
}

// GetBillSource BillSource Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetBillSource() string {
	return r._billSource
}

// SetMidPatChannel is MidPatChannel Setter
// 若无特殊物联卡传入122
func (r *AlibabaAliqinFcIotModbindAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaAliqinFcIotModbindAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}

var poolAlibabaAliqinFcIotModbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcIotModbindRequest()
	},
}

// GetAlibabaAliqinFcIotModbindRequest 从 sync.Pool 获取 AlibabaAliqinFcIotModbindAPIRequest
func GetAlibabaAliqinFcIotModbindAPIRequest() *AlibabaAliqinFcIotModbindAPIRequest {
	return poolAlibabaAliqinFcIotModbindAPIRequest.Get().(*AlibabaAliqinFcIotModbindAPIRequest)
}

// ReleaseAlibabaAliqinFcIotModbindAPIRequest 将 AlibabaAliqinFcIotModbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcIotModbindAPIRequest(v *AlibabaAliqinFcIotModbindAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcIotModbindAPIRequest.Put(v)
}
