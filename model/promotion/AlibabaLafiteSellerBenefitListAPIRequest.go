package promotion

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetApiMethodName() string {
	return "alibaba.lafite.seller.benefit.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLafiteSellerBenefitListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
