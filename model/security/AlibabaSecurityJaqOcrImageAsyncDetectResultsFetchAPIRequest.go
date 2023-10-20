package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest 聚安全获取异步图文识别结果接口 API请求
// alibaba.security.jaq.ocr.image.async.detect.results.fetch
//
// 获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
type AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest struct {
	model.Params
	// 值为图像检测接口异步调用时返回的图片task_id
	_taskIds []string
}

// NewAlibabasecurityjaqocrimageasyncdetectresultsfetchRequest 初始化AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest对象
func NewAlibabasecurityjaqocrimageasyncdetectresultsfetchRequest() *AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest {
	return &AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.ocr.image.async.detect.results.fetch"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskIds is TaskIds Setter
// 值为图像检测接口异步调用时返回的图片task_id
func (r *AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest) SetTaskIds(_taskIds []string) error {
	r._taskIds = _taskIds
	r.Set("task_ids", _taskIds)
	return nil
}

// GetTaskIds TaskIds Getter
func (r AlibabasecurityjaqocrimageasyncdetectresultsfetchAPIRequest) GetTaskIds() []string {
	return r._taskIds
}
