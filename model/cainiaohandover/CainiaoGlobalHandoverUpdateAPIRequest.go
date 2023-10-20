package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandoverupdateAPIRequest 修改交接单 API请求
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
type CainiaoglobalhandoverupdateAPIRequest struct {
	model.Params
	// 要创建交接单的小包编码集合，数量上限200
	_orderCodeList []string
	// 重量单位，克:g, 千克:kg，默认g
	_weightUnit string
	// 大包备注
	_remark string
	// 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
	_type string
	// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 用户信息
	_userInfo *UserInfoDto
	// 大包重量
	_weight int64
	// 交接单id
	_handoverOrderId int64
	// 退件信息
	_returnInfo *ReturnerDto
	// 揽收信息
	_pickupInfo *PickupDto
}

// NewCainiaoglobalhandoverupdateRequest 初始化CainiaoglobalhandoverupdateAPIRequest对象
func NewCainiaoglobalhandoverupdateRequest() *CainiaoglobalhandoverupdateAPIRequest {
	return &CainiaoglobalhandoverupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandoverupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandoverupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandoverupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCodeList is OrderCodeList Setter
// 要创建交接单的小包编码集合，数量上限200
func (r *CainiaoglobalhandoverupdateAPIRequest) SetOrderCodeList(_orderCodeList []string) error {
	r._orderCodeList = _orderCodeList
	r.Set("order_code_list", _orderCodeList)
	return nil
}

// GetOrderCodeList OrderCodeList Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetOrderCodeList() []string {
	return r._orderCodeList
}

// SetWeightUnit is WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoglobalhandoverupdateAPIRequest) SetWeightUnit(_weightUnit string) error {
	r._weightUnit = _weightUnit
	r.Set("weight_unit", _weightUnit)
	return nil
}

// GetWeightUnit WeightUnit Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetWeightUnit() string {
	return r._weightUnit
}

// SetRemark is Remark Setter
// 大包备注
func (r *CainiaoglobalhandoverupdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetRemark() string {
	return r._remark
}

// SetType is Type Setter
// 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
func (r *CainiaoglobalhandoverupdateAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetType() string {
	return r._type
}

// SetClient is Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandoverupdateAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandoverupdateAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoglobalhandoverupdateAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetWeight is Weight Setter
// 大包重量
func (r *CainiaoglobalhandoverupdateAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetHandoverOrderId is HandoverOrderId Setter
// 交接单id
func (r *CainiaoglobalhandoverupdateAPIRequest) SetHandoverOrderId(_handoverOrderId int64) error {
	r._handoverOrderId = _handoverOrderId
	r.Set("handover_order_id", _handoverOrderId)
	return nil
}

// GetHandoverOrderId HandoverOrderId Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetHandoverOrderId() int64 {
	return r._handoverOrderId
}

// SetReturnInfo is ReturnInfo Setter
// 退件信息
func (r *CainiaoglobalhandoverupdateAPIRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
	r._returnInfo = _returnInfo
	r.Set("return_info", _returnInfo)
	return nil
}

// GetReturnInfo ReturnInfo Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetReturnInfo() *ReturnerDto {
	return r._returnInfo
}

// SetPickupInfo is PickupInfo Setter
// 揽收信息
func (r *CainiaoglobalhandoverupdateAPIRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
	r._pickupInfo = _pickupInfo
	r.Set("pickup_info", _pickupInfo)
	return nil
}

// GetPickupInfo PickupInfo Getter
func (r CainiaoglobalhandoverupdateAPIRequest) GetPickupInfo() *PickupDto {
	return r._pickupInfo
}
