package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsThirdpartCompanyListAPIRequest 供应商-异云-第三方物流公司列表 API请求
// alibaba.lst.logistics.thirdpart.company.list
//
// 异地云仓发货时，需填写的第三方物流公司列表
type AlibabaLstLogisticsThirdpartCompanyListAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsCompanyQuery
}

// NewAlibabaLstLogisticsThirdpartCompanyListRequest 初始化AlibabaLstLogisticsThirdpartCompanyListAPIRequest对象
func NewAlibabaLstLogisticsThirdpartCompanyListRequest() *AlibabaLstLogisticsThirdpartCompanyListAPIRequest {
	return &AlibabaLstLogisticsThirdpartCompanyListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsThirdpartCompanyListAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.thirdpart.company.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsThirdpartCompanyListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 入参
func (r *AlibabaLstLogisticsThirdpartCompanyListAPIRequest) SetQuery(_query *LstLogisticsCompanyQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLstLogisticsThirdpartCompanyListAPIRequest) GetQuery() *LstLogisticsCompanyQuery {
	return r._query
}
