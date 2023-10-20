package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivityupdateAPIRequest 更新商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
type AlibabaretailmarketingitempoolactivityupdateAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaretailmarketingitempoolactivityupdateRequest 初始化AlibabaretailmarketingitempoolactivityupdateAPIRequest对象
func NewAlibabaretailmarketingitempoolactivityupdateRequest() *AlibabaretailmarketingitempoolactivityupdateAPIRequest {
	return &AlibabaretailmarketingitempoolactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivityupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新商品池活动参数
func (r *AlibabaretailmarketingitempoolactivityupdateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitempoolactivityupdateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
