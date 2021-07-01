package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest
更新单品买赠活动 API请求
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新 */
type AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest struct {
	model.Params
	// 更新单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivityUpdateRequest 初始化AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityUpdateRequest() *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 更新单品买赠活动参数
func (r *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}
