package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectKeywordListAPIRequest 关键词报表 API请求
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
type AlibabaScbpEffectKeywordListAPIRequest struct {
	model.Params
	// IKeywordQuery
	_p4pKeywordReportQuery *IKeywordQuery
}

// NewAlibabaScbpEffectKeywordListRequest 初始化AlibabaScbpEffectKeywordListAPIRequest对象
func NewAlibabaScbpEffectKeywordListRequest() *AlibabaScbpEffectKeywordListAPIRequest {
	return &AlibabaScbpEffectKeywordListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpEffectKeywordListAPIRequest) Reset() {
	r._p4pKeywordReportQuery = nil
	r.Params.ToZero()
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
func (r *AlibabaScbpEffectKeywordListAPIRequest) SetP4pKeywordReportQuery(_p4pKeywordReportQuery *IKeywordQuery) error {
	r._p4pKeywordReportQuery = _p4pKeywordReportQuery
	r.Set("p4p_keyword_report_query", _p4pKeywordReportQuery)
	return nil
}

// GetP4pKeywordReportQuery P4pKeywordReportQuery Getter
func (r AlibabaScbpEffectKeywordListAPIRequest) GetP4pKeywordReportQuery() *IKeywordQuery {
	return r._p4pKeywordReportQuery
}

var poolAlibabaScbpEffectKeywordListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpEffectKeywordListRequest()
	},
}

// GetAlibabaScbpEffectKeywordListRequest 从 sync.Pool 获取 AlibabaScbpEffectKeywordListAPIRequest
func GetAlibabaScbpEffectKeywordListAPIRequest() *AlibabaScbpEffectKeywordListAPIRequest {
	return poolAlibabaScbpEffectKeywordListAPIRequest.Get().(*AlibabaScbpEffectKeywordListAPIRequest)
}

// ReleaseAlibabaScbpEffectKeywordListAPIRequest 将 AlibabaScbpEffectKeywordListAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpEffectKeywordListAPIRequest(v *AlibabaScbpEffectKeywordListAPIRequest) {
	v.Reset()
	poolAlibabaScbpEffectKeywordListAPIRequest.Put(v)
}
