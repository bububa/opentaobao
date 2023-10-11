package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairQueryitemAPIRequest 换购商品查询 API请求
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabaHmMarketingItempoolStairQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaHmMarketingItempoolStairQueryitemRequest 初始化AlibabaHmMarketingItempoolStairQueryitemAPIRequest对象
func NewAlibabaHmMarketingItempoolStairQueryitemRequest() *AlibabaHmMarketingItempoolStairQueryitemAPIRequest {
	return &AlibabaHmMarketingItempoolStairQueryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.stair.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabaHmMarketingItempoolStairQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}
