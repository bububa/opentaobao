package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRcChangestatusAPIRequest 图文草稿状态更新 API请求
// alibaba.alihouse.newhome.rc.changestatus
//
// 图文草稿状态更新
type AlibabaAlihouseNewhomeRcChangestatusAPIRequest struct {
	model.Params
	// 外部图文id
	_outerId string
	// 外部楼盘id
	_outerProjectId string
	// 外部门店id
	_outerStoreId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaAlihouseNewhomeRcChangestatusRequest 初始化AlibabaAlihouseNewhomeRcChangestatusAPIRequest对象
func NewAlibabaAlihouseNewhomeRcChangestatusRequest() *AlibabaAlihouseNewhomeRcChangestatusAPIRequest {
	return &AlibabaAlihouseNewhomeRcChangestatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) Reset() {
	r._outerId = ""
	r._outerProjectId = ""
	r._outerStoreId = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.rc.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部图文id
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店id
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetStatus is Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihouseNewhomeRcChangestatusAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaAlihouseNewhomeRcChangestatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeRcChangestatusRequest()
	},
}

// GetAlibabaAlihouseNewhomeRcChangestatusRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeRcChangestatusAPIRequest
func GetAlibabaAlihouseNewhomeRcChangestatusAPIRequest() *AlibabaAlihouseNewhomeRcChangestatusAPIRequest {
	return poolAlibabaAlihouseNewhomeRcChangestatusAPIRequest.Get().(*AlibabaAlihouseNewhomeRcChangestatusAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeRcChangestatusAPIRequest 将 AlibabaAlihouseNewhomeRcChangestatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRcChangestatusAPIRequest(v *AlibabaAlihouseNewhomeRcChangestatusAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRcChangestatusAPIRequest.Put(v)
}
