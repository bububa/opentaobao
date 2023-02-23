package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverCommitAPIRequest 提交发布交接单 API请求
// cainiao.global.handover.commit
//
// 提供给ISV通过该接口提交发布交接单
type CainiaoGlobalHandoverCommitAPIRequest struct {
	model.Params
	// 要创建交接单的小包编码集合，数量上限1000
	_orderCodeList []string
	// 大包备注
	_remark string
	// 重量单位，克:g, 千克:kg，默认g
	_weightUnit string
	// 交接单类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)、self_send(自送)
	_type string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 预约交货方式（bigbag：大包预约，batch：批次预约））
	_appointmentType string
	// 国内运单号（交接单类型type为self_post(自寄)时必填）
	_domesticTrackingNo string
	// 国内物流公司名称（交接单类型type为self_post(自寄)时必填）
	_domesticLogisticsCompany string
	// 退件信息
	_returnInfo *ReturnerDto
	// 揽收信息
	_pickupInfo *PickupDto
	// 大包重量
	_weight int64
	// 交接单id
	_handoverOrderId int64
	// 用户信息
	_userInfo *UserInfoDto
	// 扩展字段
	_features *Features
	// 国内物流公司id（交接单类型type为self_post(自寄)时必填）
	_domesticLogisticsCompanyId int64
}

// NewCainiaoGlobalHandoverCommitRequest 初始化CainiaoGlobalHandoverCommitAPIRequest对象
func NewCainiaoGlobalHandoverCommitRequest() *CainiaoGlobalHandoverCommitAPIRequest {
	return &CainiaoGlobalHandoverCommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverCommitAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverCommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalHandoverCommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCodeList is OrderCodeList Setter
// 要创建交接单的小包编码集合，数量上限1000
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetOrderCodeList(_orderCodeList []string) error {
	r._orderCodeList = _orderCodeList
	r.Set("order_code_list", _orderCodeList)
	return nil
}

// GetOrderCodeList OrderCodeList Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetOrderCodeList() []string {
	return r._orderCodeList
}

// SetRemark is Remark Setter
// 大包备注
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetRemark() string {
	return r._remark
}

// SetWeightUnit is WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetWeightUnit(_weightUnit string) error {
	r._weightUnit = _weightUnit
	r.Set("weight_unit", _weightUnit)
	return nil
}

// GetWeightUnit WeightUnit Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetWeightUnit() string {
	return r._weightUnit
}

// SetType is Type Setter
// 交接单类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)、self_send(自送)
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetType() string {
	return r._type
}

// SetClient is Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetLocale() string {
	return r._locale
}

// SetAppointmentType is AppointmentType Setter
// 预约交货方式（bigbag：大包预约，batch：批次预约））
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetAppointmentType(_appointmentType string) error {
	r._appointmentType = _appointmentType
	r.Set("appointment_type", _appointmentType)
	return nil
}

// GetAppointmentType AppointmentType Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetAppointmentType() string {
	return r._appointmentType
}

// SetDomesticTrackingNo is DomesticTrackingNo Setter
// 国内运单号（交接单类型type为self_post(自寄)时必填）
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetDomesticTrackingNo(_domesticTrackingNo string) error {
	r._domesticTrackingNo = _domesticTrackingNo
	r.Set("domestic_tracking_no", _domesticTrackingNo)
	return nil
}

// GetDomesticTrackingNo DomesticTrackingNo Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetDomesticTrackingNo() string {
	return r._domesticTrackingNo
}

// SetDomesticLogisticsCompany is DomesticLogisticsCompany Setter
// 国内物流公司名称（交接单类型type为self_post(自寄)时必填）
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetDomesticLogisticsCompany(_domesticLogisticsCompany string) error {
	r._domesticLogisticsCompany = _domesticLogisticsCompany
	r.Set("domestic_logistics_company", _domesticLogisticsCompany)
	return nil
}

// GetDomesticLogisticsCompany DomesticLogisticsCompany Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetDomesticLogisticsCompany() string {
	return r._domesticLogisticsCompany
}

// SetReturnInfo is ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
	r._returnInfo = _returnInfo
	r.Set("return_info", _returnInfo)
	return nil
}

// GetReturnInfo ReturnInfo Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetReturnInfo() *ReturnerDto {
	return r._returnInfo
}

// SetPickupInfo is PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
	r._pickupInfo = _pickupInfo
	r.Set("pickup_info", _pickupInfo)
	return nil
}

// GetPickupInfo PickupInfo Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetPickupInfo() *PickupDto {
	return r._pickupInfo
}

// SetWeight is Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetHandoverOrderId is HandoverOrderId Setter
// 交接单id
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetHandoverOrderId(_handoverOrderId int64) error {
	r._handoverOrderId = _handoverOrderId
	r.Set("handover_order_id", _handoverOrderId)
	return nil
}

// GetHandoverOrderId HandoverOrderId Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetHandoverOrderId() int64 {
	return r._handoverOrderId
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetFeatures is Features Setter
// 扩展字段
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetFeatures(_features *Features) error {
	r._features = _features
	r.Set("features", _features)
	return nil
}

// GetFeatures Features Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetFeatures() *Features {
	return r._features
}

// SetDomesticLogisticsCompanyId is DomesticLogisticsCompanyId Setter
// 国内物流公司id（交接单类型type为self_post(自寄)时必填）
func (r *CainiaoGlobalHandoverCommitAPIRequest) SetDomesticLogisticsCompanyId(_domesticLogisticsCompanyId int64) error {
	r._domesticLogisticsCompanyId = _domesticLogisticsCompanyId
	r.Set("domestic_logistics_company_id", _domesticLogisticsCompanyId)
	return nil
}

// GetDomesticLogisticsCompanyId DomesticLogisticsCompanyId Getter
func (r CainiaoGlobalHandoverCommitAPIRequest) GetDomesticLogisticsCompanyId() int64 {
	return r._domesticLogisticsCompanyId
}
