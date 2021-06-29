package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片翻译任务结果 APIRequest
alibaba.seaking.imagetranslate.result

获取图片翻译任务结果
*/
type AlibabaSeakingImagetranslateResultRequest struct {
    model.Params

    // token来源站点
    tokenFrom   string 

    // 任务id
    taskId   int64 

    // 用户token
    token   string 

}

func NewAlibabaSeakingImagetranslateResultRequest() *AlibabaSeakingImagetranslateResultRequest{
    return &AlibabaSeakingImagetranslateResultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingImagetranslateResultRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate.result"
}

func (r AlibabaSeakingImagetranslateResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingImagetranslateResultRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

func (r AlibabaSeakingImagetranslateResultRequest) GetTokenFrom() string {
    return r.tokenFrom
}

func (r *AlibabaSeakingImagetranslateResultRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

func (r AlibabaSeakingImagetranslateResultRequest) GetTaskId() int64 {
    return r.taskId
}

func (r *AlibabaSeakingImagetranslateResultRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaSeakingImagetranslateResultRequest) GetToken() string {
    return r.token
}

