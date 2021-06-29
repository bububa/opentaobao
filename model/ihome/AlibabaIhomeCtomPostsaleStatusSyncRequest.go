package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
C2M售后状态同步 API请求
alibaba.ihome.ctom.postsale.status.sync

供给三维家同步定制、成品商品售后进度状态
*/
type AlibabaIhomeCtomPostsaleStatusSyncRequest struct {
    model.Params
    // 三维家服务ID
    _serviceId   string
    // 三维家售后单号ID
    _postSalesId   string
    // 三维家订单号ID
    _subOrderId   string
    // 三维家操作人部门ID
    _unitId   string
    // 三维家操作人ID
    _operatorId   string
    // 售后状态更新
    _status   string
    // 售后发起来源
    _source   string
    // 是否加急订单
    _isExpedited   string
    // 售后单更新状态原因
    _reason   string
    // 售后单结束原因
    _finishType   string
    // 客服代表ID
    _csrId   string
}

// 初始化AlibabaIhomeCtomPostsaleStatusSyncRequest对象
func NewAlibabaIhomeCtomPostsaleStatusSyncRequest() *AlibabaIhomeCtomPostsaleStatusSyncRequest{
    return &AlibabaIhomeCtomPostsaleStatusSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.postsale.status.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceId Setter
// 三维家服务ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetServiceId(_serviceId string) error {
    r._serviceId = _serviceId
    r.Set("service_id", _serviceId)
    return nil
}

// ServiceId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetServiceId() string {
    return r._serviceId
}
// PostSalesId Setter
// 三维家售后单号ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetPostSalesId(_postSalesId string) error {
    r._postSalesId = _postSalesId
    r.Set("post_sales_id", _postSalesId)
    return nil
}

// PostSalesId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetPostSalesId() string {
    return r._postSalesId
}
// SubOrderId Setter
// 三维家订单号ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetSubOrderId(_subOrderId string) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetSubOrderId() string {
    return r._subOrderId
}
// UnitId Setter
// 三维家操作人部门ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetUnitId(_unitId string) error {
    r._unitId = _unitId
    r.Set("unit_id", _unitId)
    return nil
}

// UnitId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetUnitId() string {
    return r._unitId
}
// OperatorId Setter
// 三维家操作人ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetOperatorId(_operatorId string) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetOperatorId() string {
    return r._operatorId
}
// Status Setter
// 售后状态更新
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetStatus() string {
    return r._status
}
// Source Setter
// 售后发起来源
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetSource() string {
    return r._source
}
// IsExpedited Setter
// 是否加急订单
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetIsExpedited(_isExpedited string) error {
    r._isExpedited = _isExpedited
    r.Set("is_expedited", _isExpedited)
    return nil
}

// IsExpedited Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetIsExpedited() string {
    return r._isExpedited
}
// Reason Setter
// 售后单更新状态原因
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetReason() string {
    return r._reason
}
// FinishType Setter
// 售后单结束原因
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetFinishType(_finishType string) error {
    r._finishType = _finishType
    r.Set("finish_type", _finishType)
    return nil
}

// FinishType Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetFinishType() string {
    return r._finishType
}
// CsrId Setter
// 客服代表ID
func (r *AlibabaIhomeCtomPostsaleStatusSyncRequest) SetCsrId(_csrId string) error {
    r._csrId = _csrId
    r.Set("csr_id", _csrId)
    return nil
}

// CsrId Getter
func (r AlibabaIhomeCtomPostsaleStatusSyncRequest) GetCsrId() string {
    return r._csrId
}
