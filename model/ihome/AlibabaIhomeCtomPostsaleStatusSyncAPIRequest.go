package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIhomeCtomPostsaleStatusSyncAPIRequest C2M售后状态同步 API请求
// alibaba.ihome.ctom.postsale.status.sync
//
// 供给三维家同步定制、成品商品售后进度状态
type AlibabaIhomeCtomPostsaleStatusSyncAPIRequest struct {
	model.Params
	// 三维家服务ID
	_serviceId string
	// 三维家售后单号ID
	_postSalesId string
	// 三维家订单号ID
	_subOrderId string
	// 三维家操作人部门ID
	_unitId string
	// 三维家操作人ID
	_operatorId string
	// 售后状态更新
	_status string
	// 售后发起来源
	_source string
	// 是否加急订单
	_isExpedited string
	// 售后单更新状态原因
	_reason string
	// 售后单结束原因
	_finishType string
	// 客服代表ID
	_csrId string
}

// NewAlibabaIhomeCtomPostsaleStatusSyncRequest 初始化AlibabaIhomeCtomPostsaleStatusSyncAPIRequest对象
func NewAlibabaIhomeCtomPostsaleStatusSyncRequest() *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest {
	return &AlibabaIhomeCtomPostsaleStatusSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.ihome.ctom.postsale.status.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceId Setter
// 三维家服务ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetServiceId(_serviceId string) error {
	r._serviceId = _serviceId
	r.Set("service_id", _serviceId)
	return nil
}

// Get ServiceId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetServiceId() string {
	return r._serviceId
}

// Set is PostSalesId Setter
// 三维家售后单号ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetPostSalesId(_postSalesId string) error {
	r._postSalesId = _postSalesId
	r.Set("post_sales_id", _postSalesId)
	return nil
}

// Get PostSalesId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetPostSalesId() string {
	return r._postSalesId
}

// Set is SubOrderId Setter
// 三维家订单号ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetSubOrderId(_subOrderId string) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// Get SubOrderId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetSubOrderId() string {
	return r._subOrderId
}

// Set is UnitId Setter
// 三维家操作人部门ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetUnitId(_unitId string) error {
	r._unitId = _unitId
	r.Set("unit_id", _unitId)
	return nil
}

// Get UnitId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetUnitId() string {
	return r._unitId
}

// Set is OperatorId Setter
// 三维家操作人ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// Get OperatorId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// Set is Status Setter
// 售后状态更新
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetStatus() string {
	return r._status
}

// Set is Source Setter
// 售后发起来源
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// Get Source Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetSource() string {
	return r._source
}

// Set is IsExpedited Setter
// 是否加急订单
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetIsExpedited(_isExpedited string) error {
	r._isExpedited = _isExpedited
	r.Set("is_expedited", _isExpedited)
	return nil
}

// Get IsExpedited Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetIsExpedited() string {
	return r._isExpedited
}

// Set is Reason Setter
// 售后单更新状态原因
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// Get Reason Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetReason() string {
	return r._reason
}

// Set is FinishType Setter
// 售后单结束原因
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetFinishType(_finishType string) error {
	r._finishType = _finishType
	r.Set("finish_type", _finishType)
	return nil
}

// Get FinishType Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetFinishType() string {
	return r._finishType
}

// Set is CsrId Setter
// 客服代表ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) SetCsrId(_csrId string) error {
	r._csrId = _csrId
	r.Set("csr_id", _csrId)
	return nil
}

// Get CsrId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncAPIRequest) GetCsrId() string {
	return r._csrId
}
