package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolQueryactivityAPIRequest 查找商品池活动 API请求
// alibaba.hm.marketing.itempool.queryactivity
//
// 查找商品池活动
type AlibabaHmMarketingItempoolQueryactivityAPIRequest struct {
	model.Params
	// 查询商品池活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolQueryactivityRequest 初始化AlibabaHmMarketingItempoolQueryactivityAPIRequest对象
func NewAlibabaHmMarketingItempoolQueryactivityRequest() *AlibabaHmMarketingItempoolQueryactivityAPIRequest {
	return &AlibabaHmMarketingItempoolQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询商品池活动入参
func (r *AlibabaHmMarketingItempoolQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
