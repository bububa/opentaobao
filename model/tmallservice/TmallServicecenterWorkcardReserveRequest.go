package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单预约 API请求
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

// 初始化TmallServicecenterWorkcardReserveRequest对象
func NewTmallServicecenterWorkcardReserveRequest() *TmallServicecenterWorkcardReserveRequest{
    return &TmallServicecenterWorkcardReserveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReserveRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reserve"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReserveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardReserveRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardReserveRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// ReserveTimeStart Setter
// 服务开始时间
func (r *TmallServicecenterWorkcardReserveRequest) SetReserveTimeStart(reserveTimeStart string) error {
    r.reserveTimeStart = reserveTimeStart
    r.Set("reserve_time_start", reserveTimeStart)
    return nil
}

// ReserveTimeStart Getter
func (r TmallServicecenterWorkcardReserveRequest) GetReserveTimeStart() string {
    return r.reserveTimeStart
}
// ReserveTimeEnd Setter
// 服务结束时间
func (r *TmallServicecenterWorkcardReserveRequest) SetReserveTimeEnd(reserveTimeEnd string) error {
    r.reserveTimeEnd = reserveTimeEnd
    r.Set("reserve_time_end", reserveTimeEnd)
    return nil
}

// ReserveTimeEnd Getter
func (r TmallServicecenterWorkcardReserveRequest) GetReserveTimeEnd() string {
    return r.reserveTimeEnd
}
// WorkerMobile Setter
// 工人手机号
func (r *TmallServicecenterWorkcardReserveRequest) SetWorkerMobile(workerMobile string) error {
    r.workerMobile = workerMobile
    r.Set("worker_mobile", workerMobile)
    return nil
}

// WorkerMobile Getter
func (r TmallServicecenterWorkcardReserveRequest) GetWorkerMobile() string {
    return r.workerMobile
}
// ReserveRemark Setter
// 预约备注信息
func (r *TmallServicecenterWorkcardReserveRequest) SetReserveRemark(reserveRemark string) error {
    r.reserveRemark = reserveRemark
    r.Set("reserve_remark", reserveRemark)
    return nil
}

// ReserveRemark Getter
func (r TmallServicecenterWorkcardReserveRequest) GetReserveRemark() string {
    return r.reserveRemark
}
// WorkerName Setter
// 工人姓名
func (r *TmallServicecenterWorkcardReserveRequest) SetWorkerName(workerName string) error {
    r.workerName = workerName
    r.Set("worker_name", workerName)
    return nil
}

// WorkerName Getter
func (r TmallServicecenterWorkcardReserveRequest) GetWorkerName() string {
    return r.workerName
}
// Attributes Setter
// 扩展信息
func (r *TmallServicecenterWorkcardReserveRequest) SetAttributes(attributes string) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

// Attributes Getter
func (r TmallServicecenterWorkcardReserveRequest) GetAttributes() string {
    return r.attributes
}
// Action Setter
// 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
func (r *TmallServicecenterWorkcardReserveRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

// Action Getter
func (r TmallServicecenterWorkcardReserveRequest) GetAction() string {
    return r.action
}
// FulfilType Setter
// 履约类型，上门或者到店
func (r *TmallServicecenterWorkcardReserveRequest) SetFulfilType(fulfilType string) error {
    r.fulfilType = fulfilType
    r.Set("fulfil_type", fulfilType)
    return nil
}

// FulfilType Getter
func (r TmallServicecenterWorkcardReserveRequest) GetFulfilType() string {
    return r.fulfilType
}
// ServiceStoreCode Setter
// 门店编码
func (r *TmallServicecenterWorkcardReserveRequest) SetServiceStoreCode(serviceStoreCode string) error {
    r.serviceStoreCode = serviceStoreCode
    r.Set("service_store_code", serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterWorkcardReserveRequest) GetServiceStoreCode() string {
    return r.serviceStoreCode
}
// ServiceStoreName Setter
// 门店名称
func (r *TmallServicecenterWorkcardReserveRequest) SetServiceStoreName(serviceStoreName string) error {
    r.serviceStoreName = serviceStoreName
    r.Set("service_store_name", serviceStoreName)
    return nil
}

// ServiceStoreName Getter
func (r TmallServicecenterWorkcardReserveRequest) GetServiceStoreName() string {
    return r.serviceStoreName
}
// FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardReserveRequest) SetFulfilTaskId(fulfilTaskId int64) error {
    r.fulfilTaskId = fulfilTaskId
    r.Set("fulfil_task_id", fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r TmallServicecenterWorkcardReserveRequest) GetFulfilTaskId() int64 {
    return r.fulfilTaskId
}
