package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全获取异步图文识别结果接口 API请求
alibaba.security.jaq.ocr.image.async.detect.results.fetch

获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
*/
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest struct {
    model.Params
    // 值为图像检测接口异步调用时返回的图片task_id
    _taskIds   []string
}

// 初始化AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest对象
func NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest{
    return &AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.ocr.image.async.detect.results.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskIds Setter
// 值为图像检测接口异步调用时返回的图片task_id
func (r *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) SetTaskIds(_taskIds []string) error {
    r._taskIds = _taskIds
    r.Set("task_ids", _taskIds)
    return nil
}

// TaskIds Getter
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest) GetTaskIds() []string {
    return r._taskIds
}
