package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发布交接单 API请求
cainiao.global.handover.commit

提供给ISV通过该接口提交发布交接单
*/
type CainiaoGlobalHandoverCommitRequest struct {
    model.Params
    // 用户信息
    _userInfo   *UserInfoDto
    // 大包备注
    _remark   string
    // 退件信息
    _returnInfo   *ReturnerDto
    // 揽收信息
    _pickupInfo   *PickupDto
    // 大包重量
    _weight   int64
    // 交接单id
    _handoverOrderId   int64
    // 重量单位，克:g, 千克:kg，默认g
    _weightUnit   string
    // 交接单类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)
    _type   string
    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 要创建交接单的小包编码集合，数量上限1000
    _orderCodeList   []string
    // 多语言
    _locale   string
    // 扩展字段
    _features   *Features
}

// 初始化CainiaoGlobalHandoverCommitRequest对象
func NewCainiaoGlobalHandoverCommitRequest() *CainiaoGlobalHandoverCommitRequest{
    return &CainiaoGlobalHandoverCommitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverCommitRequest) GetApiMethodName() string {
    return "cainiao.global.handover.commit"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverCommitRequest) SetUserInfo(_userInfo *UserInfoDto) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverCommitRequest) GetUserInfo() *UserInfoDto {
    return r._userInfo
}
// Remark Setter
// 大包备注
func (r *CainiaoGlobalHandoverCommitRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r CainiaoGlobalHandoverCommitRequest) GetRemark() string {
    return r._remark
}
// ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverCommitRequest) SetReturnInfo(_returnInfo *ReturnerDto) error {
    r._returnInfo = _returnInfo
    r.Set("return_info", _returnInfo)
    return nil
}

// ReturnInfo Getter
func (r CainiaoGlobalHandoverCommitRequest) GetReturnInfo() *ReturnerDto {
    return r._returnInfo
}
// PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverCommitRequest) SetPickupInfo(_pickupInfo *PickupDto) error {
    r._pickupInfo = _pickupInfo
    r.Set("pickup_info", _pickupInfo)
    return nil
}

// PickupInfo Getter
func (r CainiaoGlobalHandoverCommitRequest) GetPickupInfo() *PickupDto {
    return r._pickupInfo
}
// Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverCommitRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r CainiaoGlobalHandoverCommitRequest) GetWeight() int64 {
    return r._weight
}
// HandoverOrderId Setter
// 交接单id
func (r *CainiaoGlobalHandoverCommitRequest) SetHandoverOrderId(_handoverOrderId int64) error {
    r._handoverOrderId = _handoverOrderId
    r.Set("handover_order_id", _handoverOrderId)
    return nil
}

// HandoverOrderId Getter
func (r CainiaoGlobalHandoverCommitRequest) GetHandoverOrderId() int64 {
    return r._handoverOrderId
}
// WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverCommitRequest) SetWeightUnit(_weightUnit string) error {
    r._weightUnit = _weightUnit
    r.Set("weight_unit", _weightUnit)
    return nil
}

// WeightUnit Getter
func (r CainiaoGlobalHandoverCommitRequest) GetWeightUnit() string {
    return r._weightUnit
}
// Type Setter
// 交接单类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)
func (r *CainiaoGlobalHandoverCommitRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r CainiaoGlobalHandoverCommitRequest) GetType() string {
    return r._type
}
// Client Setter
// ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverCommitRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverCommitRequest) GetClient() string {
    return r._client
}
// OrderCodeList Setter
// 要创建交接单的小包编码集合，数量上限1000
func (r *CainiaoGlobalHandoverCommitRequest) SetOrderCodeList(_orderCodeList []string) error {
    r._orderCodeList = _orderCodeList
    r.Set("order_code_list", _orderCodeList)
    return nil
}

// OrderCodeList Getter
func (r CainiaoGlobalHandoverCommitRequest) GetOrderCodeList() []string {
    return r._orderCodeList
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverCommitRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverCommitRequest) GetLocale() string {
    return r._locale
}
// Features Setter
// 扩展字段
func (r *CainiaoGlobalHandoverCommitRequest) SetFeatures(_features *Features) error {
    r._features = _features
    r.Set("features", _features)
    return nil
}

// Features Getter
func (r CainiaoGlobalHandoverCommitRequest) GetFeatures() *Features {
    return r._features
}
