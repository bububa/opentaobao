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
    userInfo   *UserInfoDto
    // 备注
    remark   string
    // 大包重量
    weight   int64
    // 重量单位，克:g, 千克:kg，默认g
    weightUnit   string
    // 揽收信息
    pickupInfo   *PickupDto
    // 退件信息
    returnInfo   *ReturnerDto
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string
    // 需要组装大包的小包编码集合，最多限制200个小包
    orderCodeList   []string
    // 多语言
    locale   string
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
func (r *CainiaoGlobalHandoverSavedraftRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// Remark Setter
// 备注
func (r *CainiaoGlobalHandoverSavedraftRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetRemark() string {
    return r.remark
}
// Weight Setter
// 大包重量
func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetWeight() int64 {
    return r.weight
}
// WeightUnit Setter
// 重量单位，克:g, 千克:kg，默认g
func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeightUnit(weightUnit string) error {
    r.weightUnit = weightUnit
    r.Set("weight_unit", weightUnit)
    return nil
}

// WeightUnit Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetWeightUnit() string {
    return r.weightUnit
}
// PickupInfo Setter
// 揽收信息
func (r *CainiaoGlobalHandoverSavedraftRequest) SetPickupInfo(pickupInfo *PickupDto) error {
    r.pickupInfo = pickupInfo
    r.Set("pickup_info", pickupInfo)
    return nil
}

// PickupInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetPickupInfo() *PickupDto {
    return r.pickupInfo
}
// ReturnInfo Setter
// 退件信息
func (r *CainiaoGlobalHandoverSavedraftRequest) SetReturnInfo(returnInfo *ReturnerDto) error {
    r.returnInfo = returnInfo
    r.Set("return_info", returnInfo)
    return nil
}

// ReturnInfo Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetReturnInfo() *ReturnerDto {
    return r.returnInfo
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverSavedraftRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetClient() string {
    return r.client
}
// OrderCodeList Setter
// 需要组装大包的小包编码集合，最多限制200个小包
func (r *CainiaoGlobalHandoverSavedraftRequest) SetOrderCodeList(orderCodeList []string) error {
    r.orderCodeList = orderCodeList
    r.Set("order_code_list", orderCodeList)
    return nil
}

// OrderCodeList Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetOrderCodeList() []string {
    return r.orderCodeList
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverSavedraftRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverSavedraftRequest) GetLocale() string {
    return r.locale
}
