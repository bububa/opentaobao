package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySaveAPIRequest 【同城零售】单品买赠活动保存 API请求
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
type AlibabaRetailMarketingBuygiftActivitySaveAPIRequest struct {
	model.Params
	// 保存单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySaveRequest 初始化AlibabaRetailMarketingBuygiftActivitySaveAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySaveRequest() *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 保存单品买赠活动参数
func (r *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}
