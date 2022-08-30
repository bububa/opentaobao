package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLafiteSellerActivityListAPIRequest 商家自运营活动列表 API请求
// alibaba.lafite.seller.activity.list
//
// 商家查询自己配置的活动列表
type AlibabaLafiteSellerActivityListAPIRequest struct {
	model.Params
	// 请求入参
	_query *ActivityReadTopQuery
}

// NewAlibabaLafiteSellerActivityListRequest 初始化AlibabaLafiteSellerActivityListAPIRequest对象
func NewAlibabaLafiteSellerActivityListRequest() *AlibabaLafiteSellerActivityListAPIRequest {
	return &AlibabaLafiteSellerActivityListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLafiteSellerActivityListAPIRequest) GetApiMethodName() string {
	return "alibaba.lafite.seller.activity.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLafiteSellerActivityListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQuery is Query Setter
// 请求入参
func (r *AlibabaLafiteSellerActivityListAPIRequest) SetQuery(_query *ActivityReadTopQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaLafiteSellerActivityListAPIRequest) GetQuery() *ActivityReadTopQuery {
	return r._query
}
