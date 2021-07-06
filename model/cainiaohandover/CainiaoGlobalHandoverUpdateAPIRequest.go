package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverUpdateAPIRequest 修改交接单 API请求
// cainiao.global.handover.update
//
// 提供给ISV通过该接口修改交接单
type CainiaoGlobalHandoverUpdateAPIRequest struct {
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

// NewCainiaoGlobalHandoverUpdateRequest 初始化CainiaoGlobalHandoverUpdateAPIRequest对象
func NewCainiaoGlobalHandoverUpdateRequest() *CainiaoGlobalHandoverUpdateAPIRequest {
	return &CainiaoGlobalHandoverUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCodeList is OrderCodeList Setter
// 要创建交接单的小包编码集合，数量上限200
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetOrderCodeList(_orderCodeList []string) error {
	r._orderCodeList = _orderCodeList
	r.Set("order_code_list", _orderCodeList)
	return nil
}

// GetOrderCodeList OrderCodeList Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetOrderCodeList() []string {
	return r._orderCodeList
}

// SetWeightUnit is WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetWeightUnit(_weightUnit string) error {
	r._weightUnit = _weightUnit
	r.Set("weight_unit", _weightUnit)
	return nil
}

// GetWeightUnit WeightUnit Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetWeightUnit() string {
	return r._weightUnit
}

// SetRemark is Remark Setter
// 大包备注
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetRemark() string {
	return r._remark
}

// SetType is Type Setter
// 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetType() string {
	return r._type
}

// SetClient is Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetWeight is Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetHandoverOrderId is HandoverOrderId Setter
// 交接单id
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetHandoverOrderId(_handoverOrderId int64) error {
	r._handoverOrderId = _handoverOrderId
	r.Set("handover_order_id", _handoverOrderId)
	return nil
}

// GetHandoverOrderId HandoverOrderId Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetHandoverOrderId() int64 {
	return r._handoverOrderId
}

// SetReturnInfo is ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
	r._returnInfo = _returnInfo
	r.Set("return_info", _returnInfo)
	return nil
}

// GetReturnInfo ReturnInfo Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetReturnInfo() *ReturnerDto {
	return r._returnInfo
}

// SetPickupInfo is PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverUpdateAPIRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
	r._pickupInfo = _pickupInfo
	r.Set("pickup_info", _pickupInfo)
	return nil
}

// GetPickupInfo PickupInfo Getter
func (r CainiaoGlobalHandoverUpdateAPIRequest) GetPickupInfo() *PickupDto {
	return r._pickupInfo
}
