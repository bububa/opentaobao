package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetLastEffectDateAPIRequest 获取最近报表生成时间 API请求
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
type AlibabaScbpAdReportGetLastEffectDateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabaScbpAdReportGetLastEffectDateRequest 初始化AlibabaScbpAdReportGetLastEffectDateAPIRequest对象
func NewAlibabaScbpAdReportGetLastEffectDateRequest() *AlibabaScbpAdReportGetLastEffectDateAPIRequest {
	return &AlibabaScbpAdReportGetLastEffectDateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdReportGetLastEffectDateAPIRequest) Reset() {
	r._topContext = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.last.effect.date"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportGetLastEffectDateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportGetLastEffectDateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

var poolAlibabaScbpAdReportGetLastEffectDateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdReportGetLastEffectDateRequest()
	},
}

// GetAlibabaScbpAdReportGetLastEffectDateRequest 从 sync.Pool 获取 AlibabaScbpAdReportGetLastEffectDateAPIRequest
func GetAlibabaScbpAdReportGetLastEffectDateAPIRequest() *AlibabaScbpAdReportGetLastEffectDateAPIRequest {
	return poolAlibabaScbpAdReportGetLastEffectDateAPIRequest.Get().(*AlibabaScbpAdReportGetLastEffectDateAPIRequest)
}

// ReleaseAlibabaScbpAdReportGetLastEffectDateAPIRequest 将 AlibabaScbpAdReportGetLastEffectDateAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdReportGetLastEffectDateAPIRequest(v *AlibabaScbpAdReportGetLastEffectDateAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdReportGetLastEffectDateAPIRequest.Put(v)
}
