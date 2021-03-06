package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityCreateAPIRequest 创建买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.create
//
// 同城供给买赠活动创建
type AlibabaRetailMarketingBuygiftActivityCreateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivityCreateRequest 初始化AlibabaRetailMarketingBuygiftActivityCreateAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityCreateRequest() *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}
