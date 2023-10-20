package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVideoChangestatusAPIRequest 视频草稿状态更新 API请求
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
type AlibabaAlihouseNewhomeVideoChangestatusAPIRequest struct {
	model.Params
	// 外部视频id
	_outerId string
	// 外部楼盘id（不可多值）
	_outerProjectId string
	// 外部门店id
	_outerStoreId string
	// 0 失效 1 有效
	_status int64
}

// NewAlibabaAlihouseNewhomeVideoChangestatusRequest 初始化AlibabaAlihouseNewhomeVideoChangestatusAPIRequest对象
func NewAlibabaAlihouseNewhomeVideoChangestatusRequest() *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest {
	return &AlibabaAlihouseNewhomeVideoChangestatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) Reset() {
	r._outerId = ""
	r._outerProjectId = ""
	r._outerStoreId = ""
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.video.changestatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部视频id
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterProjectId is OuterProjectId Setter
// 外部楼盘id（不可多值）
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetOuterProjectId(_outerProjectId string) error {
	r._outerProjectId = _outerProjectId
	r.Set("outer_project_id", _outerProjectId)
	return nil
}

// GetOuterProjectId OuterProjectId Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetOuterProjectId() string {
	return r._outerProjectId
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店id
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}

// SetStatus is Status Setter
// 0 失效 1 有效
func (r *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaAlihouseNewhomeVideoChangestatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeVideoChangestatusRequest()
	},
}

// GetAlibabaAlihouseNewhomeVideoChangestatusRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeVideoChangestatusAPIRequest
func GetAlibabaAlihouseNewhomeVideoChangestatusAPIRequest() *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest {
	return poolAlibabaAlihouseNewhomeVideoChangestatusAPIRequest.Get().(*AlibabaAlihouseNewhomeVideoChangestatusAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeVideoChangestatusAPIRequest 将 AlibabaAlihouseNewhomeVideoChangestatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeVideoChangestatusAPIRequest(v *AlibabaAlihouseNewhomeVideoChangestatusAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeVideoChangestatusAPIRequest.Put(v)
}
