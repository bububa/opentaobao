package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
售后上门信息同步 APIRequest
alibaba.ihome.ctom.postsale.onsite.sync

用于三维家同步售后单上门人员和时间信息
*/
type AlibabaIhomeCtomPostsaleOnsiteSyncRequest struct {
    model.Params

    // 预约后展示给客户的员工联系方式
    deliveryPhone   string 

    // 提交部门ID（预约的操作人所属部门）
    unitId   string 

    // 备注
    memo   string 

    // 预约事件发生时间
    time   int64 

    // 服务ID
    serviceId   string 

    // 预约目标时间
    deliveryDate   int64 

    // 具体操作人ID（预约人）
    operatorId   string 

    // 联系人员名字
    deliveryName   string 

    // 售后单ID
    postSalesId   string 

    // 配送、安装或上门
    type   string 

    // 三维家补单ID
    additionalOrderId   string 

}

func NewAlibabaIhomeCtomPostsaleOnsiteSyncRequest() *AlibabaIhomeCtomPostsaleOnsiteSyncRequest{
    return &AlibabaIhomeCtomPostsaleOnsiteSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.postsale.onsite.sync"
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetDeliveryPhone(deliveryPhone string) error {
    r.deliveryPhone = deliveryPhone
    r.Set("delivery_phone", deliveryPhone)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetDeliveryPhone() string {
    return r.deliveryPhone
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetUnitId(unitId string) error {
    r.unitId = unitId
    r.Set("unit_id", unitId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetUnitId() string {
    return r.unitId
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetMemo() string {
    return r.memo
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetTime(time int64) error {
    r.time = time
    r.Set("time", time)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetTime() int64 {
    return r.time
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetServiceId(serviceId string) error {
    r.serviceId = serviceId
    r.Set("service_id", serviceId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetServiceId() string {
    return r.serviceId
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetDeliveryDate(deliveryDate int64) error {
    r.deliveryDate = deliveryDate
    r.Set("delivery_date", deliveryDate)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetDeliveryDate() int64 {
    return r.deliveryDate
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetOperatorId() string {
    return r.operatorId
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetDeliveryName(deliveryName string) error {
    r.deliveryName = deliveryName
    r.Set("delivery_name", deliveryName)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetDeliveryName() string {
    return r.deliveryName
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetPostSalesId(postSalesId string) error {
    r.postSalesId = postSalesId
    r.Set("post_sales_id", postSalesId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetPostSalesId() string {
    return r.postSalesId
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetType() string {
    return r.type
}

func (r *AlibabaIhomeCtomPostsaleOnsiteSyncRequest) SetAdditionalOrderId(additionalOrderId string) error {
    r.additionalOrderId = additionalOrderId
    r.Set("additional_order_id", additionalOrderId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleOnsiteSyncRequest) GetAdditionalOrderId() string {
    return r.additionalOrderId
}

