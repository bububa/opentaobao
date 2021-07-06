package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverSavedraftAPIRequest 创建交接单草稿 API请求
// cainiao.global.handover.savedraft
//
// 提供给ISV通过该接口创建交接单草稿
type CainiaoGlobalHandoverSavedraftAPIRequest struct {
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

// NewCainiaoGlobalHandoverSavedraftRequest 初始化CainiaoGlobalHandoverSavedraftAPIRequest对象
func NewCainiaoGlobalHandoverSavedraftRequest() *CainiaoGlobalHandoverSavedraftAPIRequest {
	return &CainiaoGlobalHandoverSavedraftAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.savedraft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCodeList is OrderCodeList Setter
// 需要组装大包的小包编码集合，最多限制200个小包
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetOrderCodeList(_orderCodeList []string) error {
	r._orderCodeList = _orderCodeList
	r.Set("order_code_list", _orderCodeList)
	return nil
}

// GetOrderCodeList OrderCodeList Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetOrderCodeList() []string {
	return r._orderCodeList
}

// SetRemark is Remark Setter
// 备注
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetRemark() string {
	return r._remark
}

// SetWeightUnit is WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetWeightUnit(_weightUnit string) error {
	r._weightUnit = _weightUnit
	r.Set("weight_unit", _weightUnit)
	return nil
}

// GetWeightUnit WeightUnit Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetWeightUnit() string {
	return r._weightUnit
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetLocale() string {
	return r._locale
}

// SetUserInfo is UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetUserInfo(_userInfo *UserInfoDto) error {
	r._userInfo = _userInfo
	r.Set("user_info", _userInfo)
	return nil
}

// GetUserInfo UserInfo Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetUserInfo() *UserInfoDto {
	return r._userInfo
}

// SetWeight is Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetPickupInfo is PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
	r._pickupInfo = _pickupInfo
	r.Set("pickup_info", _pickupInfo)
	return nil
}

// GetPickupInfo PickupInfo Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetPickupInfo() *PickupDto {
	return r._pickupInfo
}

// SetReturnInfo is ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverSavedraftAPIRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
	r._returnInfo = _returnInfo
	r.Set("return_info", _returnInfo)
	return nil
}

// GetReturnInfo ReturnInfo Getter
func (r CainiaoGlobalHandoverSavedraftAPIRequest) GetReturnInfo() *ReturnerDto {
	return r._returnInfo
}
