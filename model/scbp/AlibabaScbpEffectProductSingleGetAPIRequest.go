package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectProductSingleGetAPIRequest 单个产品的报表 API请求
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
type AlibabaScbpEffectProductSingleGetAPIRequest struct {
	model.Params
	// ProductQuery
	_p4pProductReportQuery *ProductQuery
}

// NewAlibabaScbpEffectProductSingleGetRequest 初始化AlibabaScbpEffectProductSingleGetAPIRequest对象
func NewAlibabaScbpEffectProductSingleGetRequest() *AlibabaScbpEffectProductSingleGetAPIRequest {
	return &AlibabaScbpEffectProductSingleGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpEffectProductSingleGetAPIRequest) Reset() {
	r._p4pProductReportQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.effect.product.single.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetP4pProductReportQuery is P4pProductReportQuery Setter
// ProductQuery
func (r *AlibabaScbpEffectProductSingleGetAPIRequest) SetP4pProductReportQuery(_p4pProductReportQuery *ProductQuery) error {
	r._p4pProductReportQuery = _p4pProductReportQuery
	r.Set("p4p_product_report_query", _p4pProductReportQuery)
	return nil
}

// GetP4pProductReportQuery P4pProductReportQuery Getter
func (r AlibabaScbpEffectProductSingleGetAPIRequest) GetP4pProductReportQuery() *ProductQuery {
	return r._p4pProductReportQuery
}

var poolAlibabaScbpEffectProductSingleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpEffectProductSingleGetRequest()
	},
}

// GetAlibabaScbpEffectProductSingleGetRequest 从 sync.Pool 获取 AlibabaScbpEffectProductSingleGetAPIRequest
func GetAlibabaScbpEffectProductSingleGetAPIRequest() *AlibabaScbpEffectProductSingleGetAPIRequest {
	return poolAlibabaScbpEffectProductSingleGetAPIRequest.Get().(*AlibabaScbpEffectProductSingleGetAPIRequest)
}

// ReleaseAlibabaScbpEffectProductSingleGetAPIRequest 将 AlibabaScbpEffectProductSingleGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpEffectProductSingleGetAPIRequest(v *AlibabaScbpEffectProductSingleGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpEffectProductSingleGetAPIRequest.Put(v)
}
