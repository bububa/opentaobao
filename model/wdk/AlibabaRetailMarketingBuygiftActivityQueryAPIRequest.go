package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityQueryAPIRequest 查询单品买赠活动【同城零售】 API请求
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
type AlibabaRetailMarketingBuygiftActivityQueryAPIRequest struct {
	model.Params
	// 买赠活动查询入参
	_param0 *BuyGiftActivityQueryRequest
}

// NewAlibabaRetailMarketingBuygiftActivityQueryRequest 初始化AlibabaRetailMarketingBuygiftActivityQueryAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityQueryRequest() *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 买赠活动查询入参
func (r *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) SetParam0(_param0 *BuyGiftActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetParam0() *BuyGiftActivityQueryRequest {
	return r._param0
}
