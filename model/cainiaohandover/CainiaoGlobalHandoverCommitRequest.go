package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发布交接单 APIRequest
cainiao.global.handover.commit

提供给ISV通过该接口提交发布交接单
*/
type CainiaoGlobalHandoverCommitRequest struct {
    model.Params

    // 用户信息
    userInfo   *UserInfoDto 

    // 大包备注
    remark   string 

    // 退件信息
    returnInfo   *ReturnerDto 

    // 揽收信息
    pickupInfo   *PickupDto 

    // 大包重量
    weight   int64 

    // 交接单id
    handoverOrderId   int64 

    // 重量单位，克:g, 千克:kg，默认g
    weightUnit   string 

    // 交接单类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)
    type   string 

    // ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string 

    // 要创建交接单的小包编码集合，数量上限1000
    orderCodeList   []string 

    // 多语言
    locale   string 

    // 扩展字段
    features   *Features 

}

func NewCainiaoGlobalHandoverCommitRequest() *CainiaoGlobalHandoverCommitRequest{
    return &CainiaoGlobalHandoverCommitRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverCommitRequest) GetApiMethodName() string {
    return "cainiao.global.handover.commit"
}

func (r CainiaoGlobalHandoverCommitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverCommitRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverCommitRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetRemark() string {
    return r.remark
}

func (r *CainiaoGlobalHandoverCommitRequest) SetReturnInfo(returnInfo *ReturnerDto) error {
    r.returnInfo = returnInfo
    r.Set("return_info", returnInfo)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetReturnInfo() *ReturnerDto {
    return r.returnInfo
}

func (r *CainiaoGlobalHandoverCommitRequest) SetPickupInfo(pickupInfo *PickupDto) error {
    r.pickupInfo = pickupInfo
    r.Set("pickup_info", pickupInfo)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetPickupInfo() *PickupDto {
    return r.pickupInfo
}

func (r *CainiaoGlobalHandoverCommitRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetWeight() int64 {
    return r.weight
}

func (r *CainiaoGlobalHandoverCommitRequest) SetHandoverOrderId(handoverOrderId int64) error {
    r.handoverOrderId = handoverOrderId
    r.Set("handover_order_id", handoverOrderId)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetHandoverOrderId() int64 {
    return r.handoverOrderId
}

func (r *CainiaoGlobalHandoverCommitRequest) SetWeightUnit(weightUnit string) error {
    r.weightUnit = weightUnit
    r.Set("weight_unit", weightUnit)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetWeightUnit() string {
    return r.weightUnit
}

func (r *CainiaoGlobalHandoverCommitRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetType() string {
    return r.type
}

func (r *CainiaoGlobalHandoverCommitRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverCommitRequest) SetOrderCodeList(orderCodeList []string) error {
    r.orderCodeList = orderCodeList
    r.Set("order_code_list", orderCodeList)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetOrderCodeList() []string {
    return r.orderCodeList
}

func (r *CainiaoGlobalHandoverCommitRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetLocale() string {
    return r.locale
}

func (r *CainiaoGlobalHandoverCommitRequest) SetFeatures(features *Features) error {
    r.features = features
    r.Set("features", features)
    return nil
}

func (r CainiaoGlobalHandoverCommitRequest) GetFeatures() *Features {
    return r.features
}

