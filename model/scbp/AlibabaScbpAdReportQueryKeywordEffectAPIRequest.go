package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportquerykeywordeffectAPIRequest 关键词报告 API请求
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
type AlibabascbpadreportquerykeywordeffectAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 请求参数
	_keywordReportOperation *KeywordReportOperationDto
}

// NewAlibabascbpadreportquerykeywordeffectRequest 初始化AlibabascbpadreportquerykeywordeffectAPIRequest对象
func NewAlibabascbpadreportquerykeywordeffectRequest() *AlibabascbpadreportquerykeywordeffectAPIRequest {
	return &AlibabascbpadreportquerykeywordeffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportquerykeywordeffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.query.keyword.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportquerykeywordeffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportquerykeywordeffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportquerykeywordeffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportquerykeywordeffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetKeywordReportOperation is KeywordReportOperation Setter
// 请求参数
func (r *AlibabascbpadreportquerykeywordeffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
	r._keywordReportOperation = _keywordReportOperation
	r.Set("keyword_report_operation", _keywordReportOperation)
	return nil
}

// GetKeywordReportOperation KeywordReportOperation Getter
func (r AlibabascbpadreportquerykeywordeffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
	return r._keywordReportOperation
}
