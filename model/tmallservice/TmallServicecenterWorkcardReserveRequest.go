package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单预约 APIRequest
tmall.servicecenter.workcard.reserve

服务工单更新通用接口
*/
type TmallServicecenterWorkcardReserveRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

    // 服务开始时间
    reserveTimeStart   string 

    // 服务结束时间
    reserveTimeEnd   string 

    // 工人手机号
    workerMobile   string 

    // 预约备注信息
    reserveRemark   string 

    // 工人姓名
    workerName   string 

    // 扩展信息
    attributes   string 

    // 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
    action   string 

    // 履约类型，上门或者到店
    fulfilType   string 

    // 门店编码
    serviceStoreCode   string 

    // 门店名称
    serviceStoreName   string 

    // 核销单id
    fulfilTaskId   int64 

}

func NewTmallServicecenterWorkcardReserveRequest() *TmallServicecenterWorkcardReserveRequest{
    return &TmallServicecenterWorkcardReserveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardReserveRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reserve"
}

func (r TmallServicecenterWorkcardReserveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardReserveRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardReserveRequest) SetReserveTimeStart(reserveTimeStart string) error {
    r.reserveTimeStart = reserveTimeStart
    r.Set("reserve_time_start", reserveTimeStart)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetReserveTimeStart() string {
    return r.reserveTimeStart
}

func (r *TmallServicecenterWorkcardReserveRequest) SetReserveTimeEnd(reserveTimeEnd string) error {
    r.reserveTimeEnd = reserveTimeEnd
    r.Set("reserve_time_end", reserveTimeEnd)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetReserveTimeEnd() string {
    return r.reserveTimeEnd
}

func (r *TmallServicecenterWorkcardReserveRequest) SetWorkerMobile(workerMobile string) error {
    r.workerMobile = workerMobile
    r.Set("worker_mobile", workerMobile)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetWorkerMobile() string {
    return r.workerMobile
}

func (r *TmallServicecenterWorkcardReserveRequest) SetReserveRemark(reserveRemark string) error {
    r.reserveRemark = reserveRemark
    r.Set("reserve_remark", reserveRemark)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetReserveRemark() string {
    return r.reserveRemark
}

func (r *TmallServicecenterWorkcardReserveRequest) SetWorkerName(workerName string) error {
    r.workerName = workerName
    r.Set("worker_name", workerName)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetWorkerName() string {
    return r.workerName
}

func (r *TmallServicecenterWorkcardReserveRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetAttributes() string {
    return r.attributes
}

func (r *TmallServicecenterWorkcardReserveRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetAction() string {
    return r.action
}

func (r *TmallServicecenterWorkcardReserveRequest) SetFulfilType(fulfilType string) error {
    r.fulfilType = fulfilType
    r.Set("fulfil_type", fulfilType)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetFulfilType() string {
    return r.fulfilType
}

func (r *TmallServicecenterWorkcardReserveRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}

func (r *TmallServicecenterWorkcardReserveRequest) SetServiceStoreName(serviceStoreName string) error {
    r.serviceStoreName = serviceStoreName
    r.Set("service_store_name", serviceStoreName)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetServiceStoreName() string {
    return r.serviceStoreName
}

func (r *TmallServicecenterWorkcardReserveRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

func (r TmallServicecenterWorkcardReserveRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}

