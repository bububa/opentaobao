package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolstairqueryitemAPIRequest 换购商品查询 API请求
// alibaba.wdk.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabawdkmarketingitempoolstairqueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabawdkmarketingitempoolstairqueryitemRequest 初始化AlibabawdkmarketingitempoolstairqueryitemAPIRequest对象
func NewAlibabawdkmarketingitempoolstairqueryitemRequest() *AlibabawdkmarketingitempoolstairqueryitemAPIRequest {
	return &AlibabawdkmarketingitempoolstairqueryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolstairqueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.stair.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolstairqueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolstairqueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabawdkmarketingitempoolstairqueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitempoolstairqueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}
