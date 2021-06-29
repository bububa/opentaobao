package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标题改写任务结果 API请求
alibaba.seaking.titlerewrite.result

获取标题改写任务结果
*/
type AlibabaSeakingTitlerewriteResultRequest struct {
    model.Params
    // token来源站点
    _tokenFrom   string
    // 任务id
    _taskId   int64
    // 用户token
    _token   string
}

// 初始化AlibabaSeakingTitlerewriteResultRequest对象
func NewAlibabaSeakingTitlerewriteResultRequest() *AlibabaSeakingTitlerewriteResultRequest{
    return &AlibabaSeakingTitlerewriteResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTitlerewriteResultRequest) GetApiMethodName() string {
    return "alibaba.seaking.titlerewrite.result"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTitlerewriteResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTitlerewriteResultRequest) SetTokenFrom(_tokenFrom string) error {
    r._tokenFrom = _tokenFrom
    r.Set("token_from", _tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetTokenFrom() string {
    return r._tokenFrom
}
// TaskId Setter
// 任务id
func (r *AlibabaSeakingTitlerewriteResultRequest) SetTaskId(_taskId int64) error {
    r._taskId = _taskId
    r.Set("task_id", _taskId)
    return nil
}

// TaskId Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetTaskId() int64 {
    return r._taskId
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteResultRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetToken() string {
    return r._token
}
