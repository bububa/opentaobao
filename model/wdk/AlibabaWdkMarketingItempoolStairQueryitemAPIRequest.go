package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolStairQueryitemAPIRequest
换购商品查询 API请求
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询 */
type AlibabaWdkMarketingItempoolStairQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaWdkMarketingItempoolStairQueryitemRequest 初始化AlibabaWdkMarketingItempoolStairQueryitemAPIRequest对象
func NewAlibabaWdkMarketingItempoolStairQueryitemRequest() *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest {
	return &AlibabaWdkMarketingItempoolStairQueryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.stair.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 换购商品查询参数
func (r *AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingItempoolStairQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}
