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
type TmallServicecenterWorkcardReserveAPIRequest struct {
    model.Params
    // 工单id
    _workcardId   int64
    // 服务开始时间
    _reserveTimeStart   string
    // 服务结束时间
    _reserveTimeEnd   string
    // 工人手机号
    _workerMobile   string
    // 预约备注信息
    _reserveRemark   string
    // 工人姓名
    _workerName   string
    // 扩展信息
    _attributes   string
    // 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
    _action   string
    // 履约类型，上门或者到店
    _fulfilType   string
    // 门店编码
    _serviceStoreCode   string
    // 门店名称
    _serviceStoreName   string
    // 核销单id
    _fulfilTaskId   int64
}

// 初始化TmallServicecenterWorkcardReserveAPIRequest对象
func NewTmallServicecenterWorkcardReserveRequest() *TmallServicecenterWorkcardReserveAPIRequest{
    return &TmallServicecenterWorkcardReserveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardReserveAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.reserve"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardReserveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// ReserveTimeStart Setter
// 服务开始时间
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetReserveTimeStart(_reserveTimeStart string) error {
    r._reserveTimeStart = _reserveTimeStart
    r.Set("reserve_time_start", _reserveTimeStart)
    return nil
}

// ReserveTimeStart Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetReserveTimeStart() string {
    return r._reserveTimeStart
}
// ReserveTimeEnd Setter
// 服务结束时间
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
    r._reserveTimeEnd = _reserveTimeEnd
    r.Set("reserve_time_end", _reserveTimeEnd)
    return nil
}

// ReserveTimeEnd Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetReserveTimeEnd() string {
    return r._reserveTimeEnd
}
// WorkerMobile Setter
// 工人手机号
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetWorkerMobile(_workerMobile string) error {
    r._workerMobile = _workerMobile
    r.Set("worker_mobile", _workerMobile)
    return nil
}

// WorkerMobile Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetWorkerMobile() string {
    return r._workerMobile
}
// ReserveRemark Setter
// 预约备注信息
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetReserveRemark(_reserveRemark string) error {
    r._reserveRemark = _reserveRemark
    r.Set("reserve_remark", _reserveRemark)
    return nil
}

// ReserveRemark Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetReserveRemark() string {
    return r._reserveRemark
}
// WorkerName Setter
// 工人姓名
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetWorkerName(_workerName string) error {
    r._workerName = _workerName
    r.Set("worker_name", _workerName)
    return nil
}

// WorkerName Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetWorkerName() string {
    return r._workerName
}
// Attributes Setter
// 扩展信息
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetAttributes(_attributes string) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetAttributes() string {
    return r._attributes
}
// Action Setter
// 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetAction() string {
    return r._action
}
// FulfilType Setter
// 履约类型，上门或者到店
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetFulfilType(_fulfilType string) error {
    r._fulfilType = _fulfilType
    r.Set("fulfil_type", _fulfilType)
    return nil
}

// FulfilType Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetFulfilType() string {
    return r._fulfilType
}
// ServiceStoreCode Setter
// 门店编码
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
    r._serviceStoreCode = _serviceStoreCode
    r.Set("service_store_code", _serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetServiceStoreCode() string {
    return r._serviceStoreCode
}
// ServiceStoreName Setter
// 门店名称
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetServiceStoreName(_serviceStoreName string) error {
    r._serviceStoreName = _serviceStoreName
    r.Set("service_store_name", _serviceStoreName)
    return nil
}

// ServiceStoreName Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetServiceStoreName() string {
    return r._serviceStoreName
}
// FulfilTaskId Setter
// 核销单id
func (r *TmallServicecenterWorkcardReserveAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
    r._fulfilTaskId = _fulfilTaskId
    r.Set("fulfil_task_id", _fulfilTaskId)
    return nil
}

// FulfilTaskId Getter
func (r TmallServicecenterWorkcardReserveAPIRequest) GetFulfilTaskId() int64 {
    return r._fulfilTaskId
}
