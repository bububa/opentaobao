package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempoolqueryitemsAPIRequest 查询商品池活动下的商品 API请求
// alibaba.wdk.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabawdkmarketingitempoolqueryitemsAPIRequest struct {
	model.Params
	// 查询入参
	_param *ActivitySkuQuery
}

// NewAlibabawdkmarketingitempoolqueryitemsRequest 初始化AlibabawdkmarketingitempoolqueryitemsAPIRequest对象
func NewAlibabawdkmarketingitempoolqueryitemsRequest() *AlibabawdkmarketingitempoolqueryitemsAPIRequest {
	return &AlibabawdkmarketingitempoolqueryitemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempoolqueryitemsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.queryitems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempoolqueryitemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempoolqueryitemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabawdkmarketingitempoolqueryitemsAPIRequest) SetParam(_param *ActivitySkuQuery) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitempoolqueryitemsAPIRequest) GetParam() *ActivitySkuQuery {
	return r._param
}
