package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
C2M售后状态同步 APIRequest
alibaba.ihome.ctom.postsale.status.sync

供给三维家同步定制、成品商品售后进度状态
*/
type AlibabaIhomeCtomPostsaleStatusSyncRequest struct {
    model.Params

    // 三维家服务ID
    serviceId   string 

    // 三维家售后单号ID
    postSalesId   string 

    // 三维家订单号ID
    subOrderId   string 

    // 三维家操作人部门ID
    unitId   string 

    // 三维家操作人ID
    operatorId   string 

    // 售后状态更新
    status   string 

    // 售后发起来源
    source   string 

    // 是否加急订单
    isExpedited   string 

    // 售后单更新状态原因
    reason   string 

    // 售后单结束原因
    finishType   string 

    // 客服代表ID
    csrId   string 

}

func NewAlibabaIhomeCtomPostsaleStatusSyncRequest() *AlibabaIhomeCtomPostsaleStatusSyncRequest{
    return &AlibabaIhomeCtomPostsaleStatusSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.postsale.status.sync"
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetServiceId(serviceId string) error {
    r.serviceId = serviceId
    r.Set("service_id", serviceId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetServiceId() string {
    return r.serviceId
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetPostSalesId(postSalesId string) error {
    r.postSalesId = postSalesId
    r.Set("post_sales_id", postSalesId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetPostSalesId() string {
    return r.postSalesId
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetSubOrderId(subOrderId string) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetSubOrderId() string {
    return r.subOrderId
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetUnitId(unitId string) error {
    r.unitId = unitId
    r.Set("unit_id", unitId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetUnitId() string {
    return r.unitId
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetOperatorId() string {
    return r.operatorId
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetSource() string {
    return r.source
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetIsExpedited(isExpedited string) error {
    r.isExpedited = isExpedited
    r.Set("is_expedited", isExpedited)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetIsExpedited() string {
    return r.isExpedited
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetFinishType(finishType string) error {
    r.finishType = finishType
    r.Set("finish_type", finishType)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetFinishType() string {
    return r.finishType
}

func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetCsrId(csrId string) error {
    r.csrId = csrId
    r.Set("csr_id", csrId)
    return nil
}

func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetCsrId() string {
    return r.csrId
}

