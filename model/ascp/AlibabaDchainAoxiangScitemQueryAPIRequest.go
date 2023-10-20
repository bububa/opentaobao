package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitemqueryAPIRequest 货品查询 API请求
// alibaba.dchain.aoxiang.scitem.query
//
// 货品查询
type AlibabadchainaoxiangscitemqueryAPIRequest struct {
	model.Params
	// 货品查询入参
	_queryScitemRequest *QueryScItemRequest
}

// NewAlibabadchainaoxiangscitemqueryRequest 初始化AlibabadchainaoxiangscitemqueryAPIRequest对象
func NewAlibabadchainaoxiangscitemqueryRequest() *AlibabadchainaoxiangscitemqueryAPIRequest {
	return &AlibabadchainaoxiangscitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangscitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangscitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangscitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryScitemRequest is QueryScitemRequest Setter
// 货品查询入参
func (r *AlibabadchainaoxiangscitemqueryAPIRequest) SetQueryScitemRequest(_queryScitemRequest *QueryScItemRequest) error {
	r._queryScitemRequest = _queryScitemRequest
	r.Set("query_scitem_request", _queryScitemRequest)
	return nil
}

// GetQueryScitemRequest QueryScitemRequest Getter
func (r AlibabadchainaoxiangscitemqueryAPIRequest) GetQueryScitemRequest() *QueryScItemRequest {
	return r._queryScitemRequest
}
