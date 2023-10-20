package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotmodbindAPIRequest 物联网绑定/换绑API API请求
// alibaba.aliqin.fc.iot.modbind
//
// 支持用户的设备的换绑和解绑操作
type AlibabaaliqinfciotmodbindAPIRequest struct {
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

// NewAlibabaaliqinfciotmodbindRequest 初始化AlibabaaliqinfciotmodbindAPIRequest对象
func NewAlibabaaliqinfciotmodbindRequest() *AlibabaaliqinfciotmodbindAPIRequest {
	return &AlibabaaliqinfciotmodbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotmodbindAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.modbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotmodbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotmodbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpionType is OpionType Setter
// chgBind：换绑；unBind：解绑
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetOpionType(_opionType string) error {
	r._opionType = _opionType
	r.Set("opion_type", _opionType)
	return nil
}

// GetOpionType OpionType Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetOpionType() string {
	return r._opionType
}

// SetBillReal is BillReal Setter
// 物联卡和iccid保持一致
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetBillReal(_billReal string) error {
	r._billReal = _billReal
	r.Set("bill_real", _billReal)
	return nil
}

// GetBillReal BillReal Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetBillReal() string {
	return r._billReal
}

// SetIccid is Iccid Setter
// iccid （20位）
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetIccid() string {
	return r._iccid
}

// SetNewimei is Newimei Setter
// 换绑的时候必传，换的新设备imei
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetNewimei(_newimei string) error {
	r._newimei = _newimei
	r.Set("newimei", _newimei)
	return nil
}

// GetNewimei Newimei Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetNewimei() string {
	return r._newimei
}

// SetImei is Imei Setter
// 目前绑定的设备imei
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetImei(_imei string) error {
	r._imei = _imei
	r.Set("imei", _imei)
	return nil
}

// GetImei Imei Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetImei() string {
	return r._imei
}

// SetBillSource is BillSource Setter
// 物联卡业务：若无特殊为ICCID
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetBillSource(_billSource string) error {
	r._billSource = _billSource
	r.Set("bill_source", _billSource)
	return nil
}

// GetBillSource BillSource Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetBillSource() string {
	return r._billSource
}

// SetMidPatChannel is MidPatChannel Setter
// 若无特殊物联卡传入122
func (r *AlibabaaliqinfciotmodbindAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaaliqinfciotmodbindAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}
