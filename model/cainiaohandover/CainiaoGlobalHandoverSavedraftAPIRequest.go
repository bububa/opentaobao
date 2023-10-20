package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandoversavedraftAPIRequest 创建交接单草稿 API请求
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
type CainiaoglobalhandoversavedraftAPIRequest struct {
	model.Params
	// 需要组装大包的小包编码集合，最多限制200个小包
	_orderCodeList []string
	// 备注
	_remark string
	// 重量单位，克:g, 千克:kg，默认g
	_weightUnit string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 用户信息
	_userInfo *UserInfoDto
	// 大包重量
	_weight int64
	// 揽收信息
	_pickupInfo *PickupDto
	// 退件信息
	_returnInfo *ReturnerDto
}

// NewCainiaoglobalhandoversavedraftRequest 初始化CainiaoglobalhandoversavedraftAPIRequest对象
func NewCainiaoglobalhandoversavedraftRequest() *CainiaoglobalhandoversavedraftAPIRequest {
	return &CainiaoglobalhandoversavedraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalhandoversavedraftAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.savedraft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalhandoversavedraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalhandoversavedraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCodeList is OrderCodeList Setter
// 需要组装大包的小包编码集合，最多限制200个小包
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetOrderCodeList(_orderCodeList []string) error {
	r._orderCodeList = _orderCodeList
	r.Set("order_code_list", _orderCodeList)
	return nil
}

// GetOrderCodeList OrderCodeList Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetOrderCodeList() []string {
	return r._orderCodeList
}

// SetRemark is Remark Setter
// 备注
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetRemark() string {
	return r._remark
}

// SetWeightUnit is WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetWeightUnit(_weightUnit string) error {
	r._weightUnit = _weightUnit
	r.Set("weight_unit", _weightUnit)
	return nil
}

// GetWeightUnit WeightUnit Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetWeightUnit() string {
	return r._weightUnit
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetWeight is Weight Setter
// 大包重量
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetPickupInfo is PickupInfo Setter
// 揽收信息
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
	r._pickupInfo = _pickupInfo
	r.Set("pickup_info", _pickupInfo)
	return nil
}

// GetPickupInfo PickupInfo Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetPickupInfo() *PickupDto {
	return r._pickupInfo
}

// SetReturnInfo is ReturnInfo Setter
// 退件信息
func (r *CainiaoglobalhandoversavedraftAPIRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
	r._returnInfo = _returnInfo
	r.Set("return_info", _returnInfo)
	return nil
}

// GetReturnInfo ReturnInfo Getter
func (r CainiaoglobalhandoversavedraftAPIRequest) GetReturnInfo() *ReturnerDto {
	return r._returnInfo
}
