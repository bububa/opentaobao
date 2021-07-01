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
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest struct {
    model.Params
    // 值为图像检测接口异步调用时返回的图片task_id
    _taskIds   []string
}

// 初始化AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest对象
func NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest{
    return &AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.ocr.image.async.detect.results.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaskIds Setter
// 值为图像检测接口异步调用时返回的图片task_id
func (r *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) SetTaskIds(_taskIds []string) error {
    r._taskIds = _taskIds
    r.Set("task_ids", _taskIds)
    return nil
}

// TaskIds Getter
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetTaskIds() []string {
    return r._taskIds
}
