package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建交接单草稿 API请求
cainiao.global.handover.savedraft

提供给ISV通过该接口创建交接单草稿
*/
type CainiaoGlobalHandoverSavedraftRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDTO
    // 备注
    _remark   string
    // 大包重量
    _weight   int64
    // 重量单位，克:g, 千克:kg，默认g
    _weightUnit   string
    // 揽收信息
    _pickupInfo   *PickupDTO
    // 退件信息
    _returnInfo   *ReturnerDTO
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 需要组装大包的小包编码集合，最多限制200个小包
    _orderCodeList   []string
    // 多语言
    _locale   string
}

// 初始化CainiaoGlobalHandoverSavedraftRequest对象
func NewCainiaoGlobalHandoverSavedraftRequest() *CainiaoGlobalHandoverSavedraftRequest{
    return &CainiaoGlobalHandoverSavedraftRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverSavedraftRequest) GetApiMethodName() string {
    return "cainiao.global.handover.savedraft"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverSavedraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverSavedraftRequest) SetUserInfo(_userInfo *UserInfoDTO) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetUserInfo() *UserInfoDTO {
    return r._userInfo
}
// Remark Setter
// 备注
func (r *CainiaoGlobalHandoverSavedraftRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetRemark() string {
    return r._remark
}
// Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetWeight() int64 {
    return r._weight
}
// WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeightUnit(_weightUnit string) error {
    r._weightUnit = _weightUnit
    r.Set("weight_unit", _weightUnit)
    return nil
}

// WeightUnit Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetWeightUnit() string {
    return r._weightUnit
}
// PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverSavedraftRequest) SetPickupInfo(_pickupInfo *PickupDTO) error {
    r._pickupInfo = _pickupInfo
    r.Set("pickup_info", _pickupInfo)
    return nil
}

// PickupInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetPickupInfo() *PickupDTO {
    return r._pickupInfo
}
// ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverSavedraftRequest) SetReturnInfo(_returnInfo *ReturnerDTO) error {
    r._returnInfo = _returnInfo
    r.Set("return_info", _returnInfo)
    return nil
}

// ReturnInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetReturnInfo() *ReturnerDTO {
    return r._returnInfo
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverSavedraftRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetClient() string {
    return r._client
}
// OrderCodeList Setter
// 需要组装大包的小包编码集合，最多限制200个小包
func (r *CainiaoGlobalHandoverSavedraftRequest) SetOrderCodeList(_orderCodeList []string) error {
    r._orderCodeList = _orderCodeList
    r.Set("order_code_list", _orderCodeList)
    return nil
}

// OrderCodeList Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetOrderCodeList() []string {
    return r._orderCodeList
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverSavedraftRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetLocale() string {
    return r._locale
}
