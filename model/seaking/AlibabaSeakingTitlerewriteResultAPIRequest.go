package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingtitlerewriteresultAPIRequest 获取标题改写任务结果 API请求
// alibaba.seaking.titlerewrite.result
//
// 获取标题改写任务结果
type AlibabaseakingtitlerewriteresultAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
	// 任务id
	_taskId int64
}

// NewAlibabaseakingtitlerewriteresultRequest 初始化AlibabaseakingtitlerewriteresultAPIRequest对象
func NewAlibabaseakingtitlerewriteresultRequest() *AlibabaseakingtitlerewriteresultAPIRequest {
	return &AlibabaseakingtitlerewriteresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.titlerewrite.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaseakingtitlerewriteresultAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaseakingtitlerewriteresultAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetToken() string {
	return r._token
}

// SetTaskId is TaskId Setter
// 任务id
func (r *AlibabaseakingtitlerewriteresultAPIRequest) SetTaskId(_taskId int64) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r AlibabaseakingtitlerewriteresultAPIRequest) GetTaskId() int64 {
	return r._taskId
}
