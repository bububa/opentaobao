package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardreserveAPIRequest 工单预约 API请求
// tmall.servicecenter.workcard.reserve
//
// 服务工单更新通用接口
type TmallservicecenterworkcardreserveAPIRequest struct {
	model.Params
	// 服务开始时间
	_reserveTimeStart string
	// 服务结束时间
	_reserveTimeEnd string
	// 工人手机号
	_workerMobile string
	// 预约备注信息
	_reserveRemark string
	// 工人姓名
	_workerName string
	// 扩展信息
	_attributes string
	// 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
	_action string
	// 履约类型，上门或者到店
	_fulfilType string
	// 门店编码
	_serviceStoreCode string
	// 门店名称
	_serviceStoreName string
	// 原因desc
	_reasonDesc string
	// 工单id
	_workcardId int64
	// 请求来源类型
	_type int64
	// 核销单id
	_fulfilTaskId int64
	// 原因code
	_reasonCode int64
}

// NewTmallservicecenterworkcardreserveRequest 初始化TmallservicecenterworkcardreserveAPIRequest对象
func NewTmallservicecenterworkcardreserveRequest() *TmallservicecenterworkcardreserveAPIRequest {
	return &TmallservicecenterworkcardreserveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardreserveAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.reserve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardreserveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardreserveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveTimeStart is ReserveTimeStart Setter
// 服务开始时间
func (r *TmallservicecenterworkcardreserveAPIRequest) SetReserveTimeStart(_reserveTimeStart string) error {
	r._reserveTimeStart = _reserveTimeStart
	r.Set("reserve_time_start", _reserveTimeStart)
	return nil
}

// GetReserveTimeStart ReserveTimeStart Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetReserveTimeStart() string {
	return r._reserveTimeStart
}

// SetReserveTimeEnd is ReserveTimeEnd Setter
// 服务结束时间
func (r *TmallservicecenterworkcardreserveAPIRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
	r._reserveTimeEnd = _reserveTimeEnd
	r.Set("reserve_time_end", _reserveTimeEnd)
	return nil
}

// GetReserveTimeEnd ReserveTimeEnd Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetReserveTimeEnd() string {
	return r._reserveTimeEnd
}

// SetWorkerMobile is WorkerMobile Setter
// 工人手机号
func (r *TmallservicecenterworkcardreserveAPIRequest) SetWorkerMobile(_workerMobile string) error {
	r._workerMobile = _workerMobile
	r.Set("worker_mobile", _workerMobile)
	return nil
}

// GetWorkerMobile WorkerMobile Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetWorkerMobile() string {
	return r._workerMobile
}

// SetReserveRemark is ReserveRemark Setter
// 预约备注信息
func (r *TmallservicecenterworkcardreserveAPIRequest) SetReserveRemark(_reserveRemark string) error {
	r._reserveRemark = _reserveRemark
	r.Set("reserve_remark", _reserveRemark)
	return nil
}

// GetReserveRemark ReserveRemark Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetReserveRemark() string {
	return r._reserveRemark
}

// SetWorkerName is WorkerName Setter
// 工人姓名
func (r *TmallservicecenterworkcardreserveAPIRequest) SetWorkerName(_workerName string) error {
	r._workerName = _workerName
	r.Set("worker_name", _workerName)
	return nil
}

// GetWorkerName WorkerName Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetWorkerName() string {
	return r._workerName
}

// SetAttributes is Attributes Setter
// 扩展信息
func (r *TmallservicecenterworkcardreserveAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetAction is Action Setter
// 存在多个不同预约节点时需要回传。用于区分具体是哪个预约节点，例如预约上门鉴定和预约上门取件
func (r *TmallservicecenterworkcardreserveAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetAction() string {
	return r._action
}

// SetFulfilType is FulfilType Setter
// 履约类型，上门或者到店
func (r *TmallservicecenterworkcardreserveAPIRequest) SetFulfilType(_fulfilType string) error {
	r._fulfilType = _fulfilType
	r.Set("fulfil_type", _fulfilType)
	return nil
}

// GetFulfilType FulfilType Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetFulfilType() string {
	return r._fulfilType
}

// SetServiceStoreCode is ServiceStoreCode Setter
// 门店编码
func (r *TmallservicecenterworkcardreserveAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
	r._serviceStoreCode = _serviceStoreCode
	r.Set("service_store_code", _serviceStoreCode)
	return nil
}

// GetServiceStoreCode ServiceStoreCode Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetServiceStoreCode() string {
	return r._serviceStoreCode
}

// SetServiceStoreName is ServiceStoreName Setter
// 门店名称
func (r *TmallservicecenterworkcardreserveAPIRequest) SetServiceStoreName(_serviceStoreName string) error {
	r._serviceStoreName = _serviceStoreName
	r.Set("service_store_name", _serviceStoreName)
	return nil
}

// GetServiceStoreName ServiceStoreName Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetServiceStoreName() string {
	return r._serviceStoreName
}

// SetReasonDesc is ReasonDesc Setter
// 原因desc
func (r *TmallservicecenterworkcardreserveAPIRequest) SetReasonDesc(_reasonDesc string) error {
	r._reasonDesc = _reasonDesc
	r.Set("reason_desc", _reasonDesc)
	return nil
}

// GetReasonDesc ReasonDesc Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetReasonDesc() string {
	return r._reasonDesc
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardreserveAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetType is Type Setter
// 请求来源类型
func (r *TmallservicecenterworkcardreserveAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetType() int64 {
	return r._type
}

// SetFulfilTaskId is FulfilTaskId Setter
// 核销单id
func (r *TmallservicecenterworkcardreserveAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}

// SetReasonCode is ReasonCode Setter
// 原因code
func (r *TmallservicecenterworkcardreserveAPIRequest) SetReasonCode(_reasonCode int64) error {
	r._reasonCode = _reasonCode
	r.Set("reason_code", _reasonCode)
	return nil
}

// GetReasonCode ReasonCode Getter
func (r TmallservicecenterworkcardreserveAPIRequest) GetReasonCode() int64 {
	return r._reasonCode
}
