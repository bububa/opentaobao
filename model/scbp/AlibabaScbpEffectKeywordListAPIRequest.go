package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectKeywordListAPIRequest 关键词报表 API请求
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
type AlibabaScbpEffectKeywordListAPIRequest struct {
	model.Params
	// IKeywordQuery
	_p4pKeywordReportQuery *IkeywordQuery
}

// NewAlibabaScbpEffectKeywordListRequest 初始化AlibabaScbpEffectKeywordListAPIRequest对象
func NewAlibabaScbpEffectKeywordListRequest() *AlibabaScbpEffectKeywordListAPIRequest {
	return &AlibabaScbpEffectKeywordListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectKeywordListAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.keyword.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectKeywordListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpEffectKeywordListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetP4pKeywordReportQuery is P4pKeywordReportQuery Setter
// IKeywordQuery
func (r *AlibabaScbpEffectKeywordListAPIRequest) SetP4pKeywordReportQuery(_p4pKeywordReportQuery *IkeywordQuery) error {
	r._p4pKeywordReportQuery = _p4pKeywordReportQuery
	r.Set("p4p_keyword_report_query", _p4pKeywordReportQuery)
	return nil
}

// GetP4pKeywordReportQuery P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordListAPIRequest) GetP4pKeywordReportQuery() *IkeywordQuery {
	return r._p4pKeywordReportQuery
}
