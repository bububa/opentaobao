package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectKeywordSingleGetAPIRequest 单个关键词效果报表 API请求
// alibaba.scbp.effect.keyword.single.get
//
// 单个关键词效果报表
type AlibabaScbpEffectKeywordSingleGetAPIRequest struct {
	model.Params
	// IKeywordQuery
	_p4pKeywordReportQuery *IKeywordQuery
}

// NewAlibabaScbpEffectKeywordSingleGetRequest 初始化AlibabaScbpEffectKeywordSingleGetAPIRequest对象
func NewAlibabaScbpEffectKeywordSingleGetRequest() *AlibabaScbpEffectKeywordSingleGetAPIRequest {
	return &AlibabaScbpEffectKeywordSingleGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpEffectKeywordSingleGetAPIRequest) Reset() {
	r._p4pKeywordReportQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.keyword.single.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetP4pKeywordReportQuery is P4pKeywordReportQuery Setter
// IKeywordQuery
func (r *AlibabaScbpEffectKeywordSingleGetAPIRequest) SetP4pKeywordReportQuery(_p4pKeywordReportQuery *IKeywordQuery) error {
	r._p4pKeywordReportQuery = _p4pKeywordReportQuery
	r.Set("p4p_keyword_report_query", _p4pKeywordReportQuery)
	return nil
}

// GetP4pKeywordReportQuery P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordSingleGetAPIRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
	return r._p4pKeywordReportQuery
}

var poolAlibabaScbpEffectKeywordSingleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpEffectKeywordSingleGetRequest()
	},
}

// GetAlibabaScbpEffectKeywordSingleGetRequest 从 sync.Pool 获取 AlibabaScbpEffectKeywordSingleGetAPIRequest
func GetAlibabaScbpEffectKeywordSingleGetAPIRequest() *AlibabaScbpEffectKeywordSingleGetAPIRequest {
	return poolAlibabaScbpEffectKeywordSingleGetAPIRequest.Get().(*AlibabaScbpEffectKeywordSingleGetAPIRequest)
}

// ReleaseAlibabaScbpEffectKeywordSingleGetAPIRequest 将 AlibabaScbpEffectKeywordSingleGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpEffectKeywordSingleGetAPIRequest(v *AlibabaScbpEffectKeywordSingleGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpEffectKeywordSingleGetAPIRequest.Put(v)
}
