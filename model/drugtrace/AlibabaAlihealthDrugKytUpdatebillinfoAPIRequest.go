package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytupdatebillinfoAPIRequest 零售端平台单据更新 API请求
// alibaba.alihealth.drug.kyt.updatebillinfo
//
// 零售端平台单据更新
type AlibabaalihealthdrugkytupdatebillinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业ID
	_entId string
	// 操作人编码
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
	// 单据编码
	_billCode string
	// 发货单位ID
	_partnerIdSend string
	// 收货单信ID
	_partnerIdRecv string
	// 详情
	_note string
}

// NewAlibabaalihealthdrugkytupdatebillinfoRequest 初始化AlibabaalihealthdrugkytupdatebillinfoAPIRequest对象
func NewAlibabaalihealthdrugkytupdatebillinfoRequest() *AlibabaalihealthdrugkytupdatebillinfoAPIRequest {
	return &AlibabaalihealthdrugkytupdatebillinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.updatebillinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntId is EntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetEntId(_entId string) error {
	r._entId = _entId
	r.Set("ent_id", _entId)
	return nil
}

// GetEntId EntId Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetEntId() string {
	return r._entId
}

// SetIcCode is IcCode Setter
// 操作人编码
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetIcCode(_icCode string) error {
	r._icCode = _icCode
	r.Set("ic_code", _icCode)
	return nil
}

// GetIcCode IcCode Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetIcCode() string {
	return r._icCode
}

// SetBillId is BillId Setter
// 单据ID
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetBillId(_billId string) error {
	r._billId = _billId
	r.Set("bill_id", _billId)
	return nil
}

// GetBillId BillId Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetBillId() string {
	return r._billId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetBillType() string {
	return r._billType
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetPartnerIdSend is PartnerIdSend Setter
// 发货单位ID
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetPartnerIdSend(_partnerIdSend string) error {
	r._partnerIdSend = _partnerIdSend
	r.Set("partner_id_send", _partnerIdSend)
	return nil
}

// GetPartnerIdSend PartnerIdSend Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetPartnerIdSend() string {
	return r._partnerIdSend
}

// SetPartnerIdRecv is PartnerIdRecv Setter
// 收货单信ID
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
	r._partnerIdRecv = _partnerIdRecv
	r.Set("partner_id_recv", _partnerIdRecv)
	return nil
}

// GetPartnerIdRecv PartnerIdRecv Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetPartnerIdRecv() string {
	return r._partnerIdRecv
}

// SetNote is Note Setter
// 详情
func (r *AlibabaalihealthdrugkytupdatebillinfoAPIRequest) SetNote(_note string) error {
	r._note = _note
	r.Set("note", _note)
	return nil
}

// GetNote Note Getter
func (r AlibabaalihealthdrugkytupdatebillinfoAPIRequest) GetNote() string {
	return r._note
}
