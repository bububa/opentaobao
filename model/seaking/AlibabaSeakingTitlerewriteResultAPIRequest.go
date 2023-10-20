package seaking

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTitlerewriteResultAPIRequest 获取标题改写任务结果 API请求
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
type AlibabaSeakingTitlerewriteResultAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
	// 任务id
	_taskId int64
}

// NewAlibabaSeakingTitlerewriteResultRequest 初始化AlibabaSeakingTitlerewriteResultAPIRequest对象
func NewAlibabaSeakingTitlerewriteResultRequest() *AlibabaSeakingTitlerewriteResultAPIRequest {
	return &AlibabaSeakingTitlerewriteResultAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSeakingTitlerewriteResultAPIRequest) Reset() {
	r._tokenFrom = ""
	r._token = ""
	r._taskId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.titlerewrite.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTitlerewriteResultAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteResultAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetToken() string {
	return r._token
}

// SetTaskId is TaskId Setter
// 任务id
func (r *AlibabaSeakingTitlerewriteResultAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r AlibabaSeakingTitlerewriteResultAPIRequest) GetTaskId() int64 {
	return r._taskId
}

var poolAlibabaSeakingTitlerewriteResultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSeakingTitlerewriteResultRequest()
	},
}

// GetAlibabaSeakingTitlerewriteResultRequest 从 sync.Pool 获取 AlibabaSeakingTitlerewriteResultAPIRequest
func GetAlibabaSeakingTitlerewriteResultAPIRequest() *AlibabaSeakingTitlerewriteResultAPIRequest {
	return poolAlibabaSeakingTitlerewriteResultAPIRequest.Get().(*AlibabaSeakingTitlerewriteResultAPIRequest)
}

// ReleaseAlibabaSeakingTitlerewriteResultAPIRequest 将 AlibabaSeakingTitlerewriteResultAPIRequest 放入 sync.Pool
func ReleaseAlibabaSeakingTitlerewriteResultAPIRequest(v *AlibabaSeakingTitlerewriteResultAPIRequest) {
	v.Reset()
	poolAlibabaSeakingTitlerewriteResultAPIRequest.Put(v)
}
