package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交接单 API请求
cainiao.global.handover.update

提供给ISV通过该接口修改交接单
*/
type CainiaoGlobalHandoverUpdateRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDTO
    // 重量单位，克:g, 千克:kg，默认g
    _weightUnit   string
    // 大包重量
    _weight   int64
    // 交接单id
    _handoverOrderId   int64
    // 大包备注
    _remark   string
    // 退件信息
    _returnInfo   *ReturnerDTO
    // 揽收信息
    _pickupInfo   *PickupDTO
    // 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
    _type   string
    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 要创建交接单的小包编码集合，数量上限200
    _orderCodeList   []string
    // 多语言
    _locale   string
}

// 初始化CainiaoGlobalHandoverUpdateRequest对象
func NewCainiaoGlobalHandoverUpdateRequest() *CainiaoGlobalHandoverUpdateRequest{
    return &CainiaoGlobalHandoverUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverUpdateRequest) GetApiMethodName() string {
    return "cainiao.global.handover.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverUpdateRequest) SetUserInfo(_userInfo *UserInfoDTO) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetUserInfo() *UserInfoDTO {
    return r._userInfo
}
// WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverUpdateRequest) SetWeightUnit(_weightUnit string) error {
    r._weightUnit = _weightUnit
    r.Set("weight_unit", _weightUnit)
    return nil
}

// WeightUnit Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetWeightUnit() string {
    return r._weightUnit
}
// Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverUpdateRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetWeight() int64 {
    return r._weight
}
// HandoverOrderId Setter
// 交接单id
func (r *CainiaoGlobalHandoverUpdateRequest) SetHandoverOrderId(_handoverOrderId int64) error {
    r._handoverOrderId = _handoverOrderId
    r.Set("handover_order_id", _handoverOrderId)
    return nil
}

// HandoverOrderId Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetHandoverOrderId() int64 {
    return r._handoverOrderId
}
// Remark Setter
// 大包备注
func (r *CainiaoGlobalHandoverUpdateRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetRemark() string {
    return r._remark
}
// ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverUpdateRequest) SetReturnInfo(_returnInfo *ReturnerDTO) error {
    r._returnInfo = _returnInfo
    r.Set("return_info", _returnInfo)
    return nil
}

// ReturnInfo Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetReturnInfo() *ReturnerDTO {
    return r._returnInfo
}
// PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverUpdateRequest) SetPickupInfo(_pickupInfo *PickupDTO) error {
    r._pickupInfo = _pickupInfo
    r.Set("pickup_info", _pickupInfo)
    return nil
}

// PickupInfo Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetPickupInfo() *PickupDTO {
    return r._pickupInfo
}
// Type Setter
// 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
func (r *CainiaoGlobalHandoverUpdateRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetType() string {
    return r._type
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverUpdateRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetClient() string {
    return r._client
}
// OrderCodeList Setter
// 要创建交接单的小包编码集合，数量上限200
func (r *CainiaoGlobalHandoverUpdateRequest) SetOrderCodeList(_orderCodeList []string) error {
    r._orderCodeList = _orderCodeList
    r.Set("order_code_list", _orderCodeList)
    return nil
}

// OrderCodeList Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetOrderCodeList() []string {
    return r._orderCodeList
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverUpdateRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverUpdateRequest) GetLocale() string {
    return r._locale
}
