package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标题改写任务结果 APIRequest
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

func NewAlibabaSeakingTitlerewriteResultRequest() *AlibabaSeakingTitlerewriteResultRequest{
    return &AlibabaSeakingTitlerewriteResultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingTitlerewriteResultRequest) GetApiMethodName() string {
    return "alibaba.seaking.titlerewrite.result"
}

func (r AlibabaSeakingTitlerewriteResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingTitlerewriteResultRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

func (r AlibabaSeakingTitlerewriteResultRequest) GetTokenFrom() string {
    return r.tokenFrom
}

func (r *AlibabaSeakingTitlerewriteResultRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r AlibabaSeakingTitlerewriteResultRequest) GetTaskId() int64 {
    return r.taskId
}

func (r *AlibabaSeakingTitlerewriteResultRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaSeakingTitlerewriteResultRequest) GetToken() string {
    return r.token
}

