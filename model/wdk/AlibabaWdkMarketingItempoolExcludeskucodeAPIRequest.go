package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest
商品池排除商品【品类优惠使用】 API请求
alibaba.wdk.marketing.itempool.excludeskucode

品类优惠新增排除池 */
type AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolExcludeskucodeRequest 初始化AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest对象
func NewAlibabaWdkMarketingItempoolExcludeskucodeRequest() *AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest {
	return &AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.excludeskucode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// Set is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingItempoolExcludeskucodeAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
