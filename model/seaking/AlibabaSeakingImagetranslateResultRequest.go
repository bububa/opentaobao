package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片翻译任务结果 API请求
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

// 初始化AlibabaSeakingImagetranslateResultRequest对象
func NewAlibabaSeakingImagetranslateResultRequest() *AlibabaSeakingImagetranslateResultRequest{
    return &AlibabaSeakingImagetranslateResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateResultRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate.result"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingImagetranslateResultRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingImagetranslateResultRequest) GetTokenFrom() string {
    return r.tokenFrom
}
// TaskId Setter
// 任务id
func (r *AlibabaSeakingImagetranslateResultRequest) SetTaskId(taskId int64) error {
    r.taskId = taskId
    r.Set("task_id", taskId)
    return nil
}

// TaskId Getter
func (r AlibabaSeakingImagetranslateResultRequest) GetTaskId() int64 {
    return r.taskId
}
// Token Setter
// 用户token
func (r *AlibabaSeakingImagetranslateResultRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSeakingImagetranslateResultRequest) GetToken() string {
    return r.token
}
