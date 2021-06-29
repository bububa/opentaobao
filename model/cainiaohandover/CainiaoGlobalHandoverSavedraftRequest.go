package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建交接单草稿 APIRequest
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

func NewCainiaoGlobalHandoverSavedraftRequest() *CainiaoGlobalHandoverSavedraftRequest{
    return &CainiaoGlobalHandoverSavedraftRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetApiMethodName() string {
    return "cainiao.global.handover.savedraft"
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverSavedraftRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetRemark() string {
    return r.remark
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetWeight() int64 {
    return r.weight
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetWeightUnit(weightUnit string) error {
    r.weightUnit = weightUnit
    r.Set("weight_unit", weightUnit)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetWeightUnit() string {
    return r.weightUnit
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetPickupInfo(pickupInfo *PickupDto) error {
    r.pickupInfo = pickupInfo
    r.Set("pickup_info", pickupInfo)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetPickupInfo() *PickupDto {
    return r.pickupInfo
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetReturnInfo(returnInfo *ReturnerDto) error {
    r.returnInfo = returnInfo
    r.Set("return_info", returnInfo)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetReturnInfo() *ReturnerDto {
    return r.returnInfo
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetOrderCodeList(orderCodeList []string) error {
    r.orderCodeList = orderCodeList
    r.Set("order_code_list", orderCodeList)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetOrderCodeList() []string {
    return r.orderCodeList
}

func (r *CainiaoGlobalHandoverSavedraftRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverSavedraftRequest) GetLocale() string {
    return r.locale
}

