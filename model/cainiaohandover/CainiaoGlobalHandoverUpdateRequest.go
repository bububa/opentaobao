package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交接单 APIRequest
cainiao.global.handover.update

提供给ISV通过该接口修改交接单
*/
type CainiaoGlobalHandoverUpdateRequest struct {
    model.Params

    // 用户信息
    userInfo   *UserInfoDto 

    // 重量单位，克:g, 千克:kg，默认g
    weightUnit   string 

    // 大包重量
    weight   int64 

    // 交接单id
    handoverOrderId   int64 

    // 大包备注
    remark   string 

    // 退件信息
    returnInfo   *ReturnerDto 

    // 揽收信息
    pickupInfo   *PickupDto 

    // 交接单类型，菜鸟揽收(cainiao_pickup)或自寄(self_post)，默认菜鸟揽收
    type   string 

    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string 

    // 要创建交接单的小包编码集合，数量上限200
    orderCodeList   []string 

    // 多语言
    locale   string 

}

func NewCainiaoGlobalHandoverUpdateRequest() *CainiaoGlobalHandoverUpdateRequest{
    return &CainiaoGlobalHandoverUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverUpdateRequest) GetApiMethodName() string {
    return "cainiao.global.handover.update"
}

func (r CainiaoGlobalHandoverUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverUpdateRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetWeightUnit(weightUnit string) error {
    r.weightUnit = weightUnit
    r.Set("weight_unit", weightUnit)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetWeightUnit() string {
    return r.weightUnit
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetWeight() int64 {
    return r.weight
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetHandoverOrderId(handoverOrderId int64) error {
    r.handoverOrderId = handoverOrderId
    r.Set("handover_order_id", handoverOrderId)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetHandoverOrderId() int64 {
    return r.handoverOrderId
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetRemark() string {
    return r.remark
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetReturnInfo(returnInfo *ReturnerDto) error {
    r.returnInfo = returnInfo
    r.Set("return_info", returnInfo)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetReturnInfo() *ReturnerDto {
    return r.returnInfo
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetPickupInfo(pickupInfo *PickupDto) error {
    r.pickupInfo = pickupInfo
    r.Set("pickup_info", pickupInfo)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetPickupInfo() *PickupDto {
    return r.pickupInfo
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetType() string {
    return r.type
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetOrderCodeList(orderCodeList []string) error {
    r.orderCodeList = orderCodeList
    r.Set("order_code_list", orderCodeList)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetOrderCodeList() []string {
    return r.orderCodeList
}

func (r *CainiaoGlobalHandoverUpdateRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverUpdateRequest) GetLocale() string {
    return r.locale
}

