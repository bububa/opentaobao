package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest wes零头出入库单据上传 API请求
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
type AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest struct {
	model.Params
	// 企业ref_ent_id
	_refEntId string
	// 服务时间
	_licenseToken string
	// 零头入库：106；零头出库：210
	_billType string
	// 单据编号
	_billCode string
	// 单据时间:yyyy-MM-dd HH:mm:ss
	_billTime string
	// 发货企业【注意：该入参是ref_ent_id,并不是ent_id】
	_fromRefUserId string
	// 收货企业【注意：该入参是ref_ent_id,并不是ent_id】
	_toRefUserId string
	// 委托企业【注意：该入参是ref_ent_id,并不是ent_id】
	_assRefEntId string
	// 配送企业【注意：该入参是ref_ent_id,并不是ent_id】
	_disRefEntId string
	// 药品ID
	_drugEntBaseInfoId string
	// 生产日期：yyyy-MM-dd
	_produceDate string
	// 有效期:yyyyMMdd
	_expireDate string
	// 生产批次
	_produceBatchNo string
	// 药品数量
	_inputAmount string
}

// NewAlibabaAlihealthDrugKytWesRemnantbillUploadRequest 初始化AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest对象
func NewAlibabaAlihealthDrugKytWesRemnantbillUploadRequest() *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest {
	return &AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.remnantbill.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBillType is BillType Setter
// 零头入库：106；零头出库：210
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetBillType() string {
	return r._billType
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据时间:yyyy-MM-dd HH:mm:ss
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业【注意：该入参是ref_ent_id,并不是ent_id】
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业【注意：该入参是ref_ent_id,并不是ent_id】
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}

// SetAssRefEntId is AssRefEntId Setter
// 委托企业【注意：该入参是ref_ent_id,并不是ent_id】
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetAssRefEntId(_assRefEntId string) error {
	r._assRefEntId = _assRefEntId
	r.Set("ass_ref_ent_id", _assRefEntId)
	return nil
}

// GetAssRefEntId AssRefEntId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetAssRefEntId() string {
	return r._assRefEntId
}

// SetDisRefEntId is DisRefEntId Setter
// 配送企业【注意：该入参是ref_ent_id,并不是ent_id】
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetDisRefEntId(_disRefEntId string) error {
	r._disRefEntId = _disRefEntId
	r.Set("dis_ref_ent_id", _disRefEntId)
	return nil
}

// GetDisRefEntId DisRefEntId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetDisRefEntId() string {
	return r._disRefEntId
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetProduceDate is ProduceDate Setter
// 生产日期：yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetProduceDate(_produceDate string) error {
	r._produceDate = _produceDate
	r.Set("produce_date", _produceDate)
	return nil
}

// GetProduceDate ProduceDate Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetProduceDate() string {
	return r._produceDate
}

// SetExpireDate is ExpireDate Setter
// 有效期:yyyyMMdd
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetExpireDate(_expireDate string) error {
	r._expireDate = _expireDate
	r.Set("expire_date", _expireDate)
	return nil
}

// GetExpireDate ExpireDate Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetExpireDate() string {
	return r._expireDate
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 生产批次
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetInputAmount is InputAmount Setter
// 药品数量
func (r *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) SetInputAmount(_inputAmount string) error {
	r._inputAmount = _inputAmount
	r.Set("input_amount", _inputAmount)
	return nil
}

// GetInputAmount InputAmount Getter
func (r AlibabaAlihealthDrugKytWesRemnantbillUploadAPIRequest) GetInputAmount() string {
	return r._inputAmount
}
