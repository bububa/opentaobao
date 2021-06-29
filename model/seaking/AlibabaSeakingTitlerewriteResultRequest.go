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
    tokenFrom   string
    // 任务id
    taskId   int64
    // 用户token
    token   string
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
func (r *AlibabaSeakingTitlerewriteResultRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetTokenFrom() string {
    return r.tokenFrom
}
// TaskId Setter
// 任务id
func (r *AlibabaSeakingTitlerewriteResultRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetTaskId() int64 {
    return r.taskId
}
// Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteResultRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSeakingTitlerewriteResultRequest) GetToken() string {
    return r.token
}
