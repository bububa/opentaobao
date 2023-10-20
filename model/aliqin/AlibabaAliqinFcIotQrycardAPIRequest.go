package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotqrycardAPIRequest 查询终端信息 API请求
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
type AlibabaaliqinfciotqrycardAPIRequest struct {
	model.Params
	// 外部计费来源
	_billSource string
	// 外部计费号
	_billReal string
	// ICCID
	_iccid string
}

// NewAlibabaaliqinfciotqrycardRequest 初始化AlibabaaliqinfciotqrycardAPIRequest对象
func NewAlibabaaliqinfciotqrycardRequest() *AlibabaaliqinfciotqrycardAPIRequest {
	return &AlibabaaliqinfciotqrycardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotqrycardAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.qrycard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotqrycardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotqrycardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillSource is BillSource Setter
// 外部计费来源
func (r *AlibabaaliqinfciotqrycardAPIRequest) SetBillSource(_billSource string) error {
	r._billSource = _billSource
	r.Set("bill_source", _billSource)
	return nil
}

// GetBillSource BillSource Getter
func (r AlibabaaliqinfciotqrycardAPIRequest) GetBillSource() string {
	return r._billSource
}

// SetBillReal is BillReal Setter
// 外部计费号
func (r *AlibabaaliqinfciotqrycardAPIRequest) SetBillReal(_billReal string) error {
	r._billReal = _billReal
	r.Set("bill_real", _billReal)
	return nil
}

// GetBillReal BillReal Getter
func (r AlibabaaliqinfciotqrycardAPIRequest) GetBillReal() string {
	return r._billReal
}

// SetIccid is Iccid Setter
// ICCID
func (r *AlibabaaliqinfciotqrycardAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaaliqinfciotqrycardAPIRequest) GetIccid() string {
	return r._iccid
}
