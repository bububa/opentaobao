package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadreportquerysinglekeywordeffectAPIRequest 单个关键词报告 API请求
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
type AlibabascbpadreportquerysinglekeywordeffectAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 返回详情
	_keywordReportOperation *KeywordReportOperationDto
}

// NewAlibabascbpadreportquerysinglekeywordeffectRequest 初始化AlibabascbpadreportquerysinglekeywordeffectAPIRequest对象
func NewAlibabascbpadreportquerysinglekeywordeffectRequest() *AlibabascbpadreportquerysinglekeywordeffectAPIRequest {
	return &AlibabascbpadreportquerysinglekeywordeffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadreportquerysinglekeywordeffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadreportquerysinglekeywordeffectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadreportquerysinglekeywordeffectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadreportquerysinglekeywordeffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadreportquerysinglekeywordeffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetKeywordReportOperation is KeywordReportOperation Setter
// 返回详情
func (r *AlibabascbpadreportquerysinglekeywordeffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
	r._keywordReportOperation = _keywordReportOperation
	r.Set("keyword_report_operation", _keywordReportOperation)
	return nil
}

// GetKeywordReportOperation KeywordReportOperation Getter
func (r AlibabascbpadreportquerysinglekeywordeffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
	return r._keywordReportOperation
}
