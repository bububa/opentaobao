package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItempoolActivityUpdateAPIRequest
更新商品池活动【同城零售】 API请求
alibaba.retail.marketing.itempool.activity.update

同城零售商品池活动更新 */
type AlibabaRetailMarketingItempoolActivityUpdateAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivityUpdateRequest 初始化AlibabaRetailMarketingItempoolActivityUpdateAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityUpdateRequest() *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 更新商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
