package security

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest 聚安全获取异步图文识别结果接口 API请求
// alibaba.security.jaq.ocr.image.async.detect.results.fetch
//
// 获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest struct {
	model.Params
	// 值为图像检测接口异步调用时返回的图片task_id
	_taskIds []string
}

// NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest 初始化AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest对象
func NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest {
	return &AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) Reset() {
	r._taskIds = r._taskIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.ocr.image.async.detect.results.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskIds is TaskIds Setter
// 值为图像检测接口异步调用时返回的图片task_id
func (r *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) SetTaskIds(_taskIds []string) error {
	r._taskIds = _taskIds
	r.Set("task_ids", _taskIds)
	return nil
}

// GetTaskIds TaskIds Getter
func (r AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) GetTaskIds() []string {
	return r._taskIds
}

var poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest()
	},
}

// GetAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchRequest 从 sync.Pool 获取 AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest
func GetAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest() *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest {
	return poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest.Get().(*AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest)
}

// ReleaseAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest 将 AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest 放入 sync.Pool
func ReleaseAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest(v *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest) {
	v.Reset()
	poolAlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIRequest.Put(v)
}
