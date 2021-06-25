package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全获取异步图文识别结果接口 APIRequest
alibaba.security.jaq.ocr.image.async.detect.results.fetch

获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
*/
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest struct {
    model.Params

    // 值为图像检测接口异步调用时返回的图片task_id
    taskIds   []String 

}

func NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest{
    return &AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.ocr.image.async.detect.results.fetch"
}

func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) SetTaskIds(taskIds []String) error {
    r.taskIds = taskIds
    r.Set("task_ids", taskIds)
    return nil
}

func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetTaskIds() []String {
    return r.taskIds
}

