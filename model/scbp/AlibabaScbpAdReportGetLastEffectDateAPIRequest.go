package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportgetlasteffectdateAPIRequest 获取最近报表生成时间 API请求
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
type AlibabascbpadreportgetlasteffectdateAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
}

// NewAlibabascbpadreportgetlasteffectdateRequest 初始化AlibabascbpadreportgetlasteffectdateAPIRequest对象
func NewAlibabascbpadreportgetlasteffectdateRequest() *AlibabascbpadreportgetlasteffectdateAPIRequest {
	return &AlibabascbpadreportgetlasteffectdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportgetlasteffectdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.get.last.effect.date"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportgetlasteffectdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportgetlasteffectdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportgetlasteffectdateAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportgetlasteffectdateAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}
