package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest 单个关键词报告 API请求
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest struct {
	model.Params
	// 用户信息
	_topContext *TopContextDto
	// 返回详情
	_keywordReportOperation *KeywordReportOperationDto
}

// NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest 初始化AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest对象
func NewAlibabaScbpAdReportQuerySingleKeywordEffectRequest() *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest {
	return &AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.report.query.single.keyword.effect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetKeywordReportOperation is KeywordReportOperation Setter
// 返回详情
func (r *AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) SetKeywordReportOperation(_keywordReportOperation *KeywordReportOperationDto) error {
	r._keywordReportOperation = _keywordReportOperation
	r.Set("keyword_report_operation", _keywordReportOperation)
	return nil
}

// GetKeywordReportOperation KeywordReportOperation Getter
func (r AlibabaScbpAdReportQuerySingleKeywordEffectAPIRequest) GetKeywordReportOperation() *KeywordReportOperationDto {
	return r._keywordReportOperation
}
