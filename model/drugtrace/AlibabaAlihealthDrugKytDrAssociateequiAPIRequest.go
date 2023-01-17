package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrAssociateequiAPIRequest 疫苗单据与设备绑定 API请求
// alibaba.alihealth.drug.kyt.dr.associateequi
//
// 疫苗单据与设备绑定
type AlibabaAlihealthDrugKytDrAssociateequiAPIRequest struct {
	model.Params
	// 企业refentid
	_refEntId string
	// 单据编号，多个用逗号分隔
	_billCodes string
	// 设备ID
	_vaEquipmentId string
}

// NewAlibabaAlihealthDrugKytDrAssociateequiRequest 初始化AlibabaAlihealthDrugKytDrAssociateequiAPIRequest对象
func NewAlibabaAlihealthDrugKytDrAssociateequiRequest() *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest {
	return &AlibabaAlihealthDrugKytDrAssociateequiAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.associateequi"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refentid
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCodes is BillCodes Setter
// 单据编号，多个用逗号分隔
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetBillCodes(_billCodes string) error {
	r._billCodes = _billCodes
	r.Set("bill_codes", _billCodes)
	return nil
}

// GetBillCodes BillCodes Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetBillCodes() string {
	return r._billCodes
}

// SetVaEquipmentId is VaEquipmentId Setter
// 设备ID
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetVaEquipmentId(_vaEquipmentId string) error {
	r._vaEquipmentId = _vaEquipmentId
	r.Set("va_equipment_id", _vaEquipmentId)
	return nil
}

// GetVaEquipmentId VaEquipmentId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetVaEquipmentId() string {
	return r._vaEquipmentId
}
