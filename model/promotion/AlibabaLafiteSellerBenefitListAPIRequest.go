package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLafiteSellerBenefitListAPIRequest 商家自运营权益列表 API请求
// alibaba.lafite.seller.benefit.list
//
// 小程序isv可使用该接口获取权益列表
type AlibabaLafiteSellerBenefitListAPIRequest struct {
	model.Params
	// 查询参数
	_benefitReadTopQuery *BenefitReadTopQuery
}

// NewAlibabaLafiteSellerBenefitListRequest 初始化AlibabaLafiteSellerBenefitListAPIRequest对象
func NewAlibabaLafiteSellerBenefitListRequest() *AlibabaLafiteSellerBenefitListAPIRequest {
	return &AlibabaLafiteSellerBenefitListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLafiteSellerBenefitListAPIRequest) Reset() {
	r._benefitReadTopQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetApiMethodName() string {
	return "alibaba.lafite.seller.benefit.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBenefitReadTopQuery is BenefitReadTopQuery Setter
// 查询参数
func (r *AlibabaLafiteSellerBenefitListAPIRequest) SetBenefitReadTopQuery(_benefitReadTopQuery *BenefitReadTopQuery) error {
	r._benefitReadTopQuery = _benefitReadTopQuery
	r.Set("benefit_read_top_query", _benefitReadTopQuery)
	return nil
}

// GetBenefitReadTopQuery BenefitReadTopQuery Getter
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetBenefitReadTopQuery() *BenefitReadTopQuery {
	return r._benefitReadTopQuery
}

var poolAlibabaLafiteSellerBenefitListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLafiteSellerBenefitListRequest()
	},
}

// GetAlibabaLafiteSellerBenefitListRequest 从 sync.Pool 获取 AlibabaLafiteSellerBenefitListAPIRequest
func GetAlibabaLafiteSellerBenefitListAPIRequest() *AlibabaLafiteSellerBenefitListAPIRequest {
	return poolAlibabaLafiteSellerBenefitListAPIRequest.Get().(*AlibabaLafiteSellerBenefitListAPIRequest)
}

// ReleaseAlibabaLafiteSellerBenefitListAPIRequest 将 AlibabaLafiteSellerBenefitListAPIRequest 放入 sync.Pool
func ReleaseAlibabaLafiteSellerBenefitListAPIRequest(v *AlibabaLafiteSellerBenefitListAPIRequest) {
	v.Reset()
	poolAlibabaLafiteSellerBenefitListAPIRequest.Put(v)
}
