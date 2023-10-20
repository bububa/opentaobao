package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixSeatTokenQueryAPIRequest 分销商选座获取qtoken API请求
// alibaba.damai.maitix.seat.token.query
//
// 选座分销，分销商查询token
type AlibabaDamaiMaitixSeatTokenQueryAPIRequest struct {
	model.Params
	// 必填-选座结束跳转回去的url,这是渠道方自己的url地址,用于接收选座后的座位信息参数
	_callbackUrl string
	// 会话ID，保证一次选座会话,建议使用 appKey+随机串 生成 ；注意：同一个场次下的会话ID不能重复
	_requestId string
	// 场次ID-必填
	_performId int64
	// 项目ID-必填
	_projectId int64
}

// NewAlibabaDamaiMaitixSeatTokenQueryRequest 初始化AlibabaDamaiMaitixSeatTokenQueryAPIRequest对象
func NewAlibabaDamaiMaitixSeatTokenQueryRequest() *AlibabaDamaiMaitixSeatTokenQueryAPIRequest {
	return &AlibabaDamaiMaitixSeatTokenQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) Reset() {
	r._callbackUrl = ""
	r._requestId = ""
	r._performId = 0
	r._projectId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.seat.token.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackUrl is CallbackUrl Setter
// 必填-选座结束跳转回去的url,这是渠道方自己的url地址,用于接收选座后的座位信息参数
func (r *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetRequestId is RequestId Setter
// 会话ID，保证一次选座会话,建议使用 appKey+随机串 生成 ；注意：同一个场次下的会话ID不能重复
func (r *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetPerformId is PerformId Setter
// 场次ID-必填
func (r *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) SetPerformId(_performId int64) error {
	r._performId = _performId
	r.Set("perform_id", _performId)
	return nil
}

// GetPerformId PerformId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetPerformId() int64 {
	return r._performId
}

// SetProjectId is ProjectId Setter
// 项目ID-必填
func (r *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) SetProjectId(_projectId int64) error {
	r._projectId = _projectId
	r.Set("project_id", _projectId)
	return nil
}

// GetProjectId ProjectId Getter
func (r AlibabaDamaiMaitixSeatTokenQueryAPIRequest) GetProjectId() int64 {
	return r._projectId
}

var poolAlibabaDamaiMaitixSeatTokenQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixSeatTokenQueryRequest()
	},
}

// GetAlibabaDamaiMaitixSeatTokenQueryRequest 从 sync.Pool 获取 AlibabaDamaiMaitixSeatTokenQueryAPIRequest
func GetAlibabaDamaiMaitixSeatTokenQueryAPIRequest() *AlibabaDamaiMaitixSeatTokenQueryAPIRequest {
	return poolAlibabaDamaiMaitixSeatTokenQueryAPIRequest.Get().(*AlibabaDamaiMaitixSeatTokenQueryAPIRequest)
}

// ReleaseAlibabaDamaiMaitixSeatTokenQueryAPIRequest 将 AlibabaDamaiMaitixSeatTokenQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixSeatTokenQueryAPIRequest(v *AlibabaDamaiMaitixSeatTokenQueryAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixSeatTokenQueryAPIRequest.Put(v)
}
