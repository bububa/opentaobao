package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivitycreateAPIRequest 创建商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
type AlibabaretailmarketingitempoolactivitycreateAPIRequest struct {
	model.Params
	// 创建商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaretailmarketingitempoolactivitycreateRequest 初始化AlibabaretailmarketingitempoolactivitycreateAPIRequest对象
func NewAlibabaretailmarketingitempoolactivitycreateRequest() *AlibabaretailmarketingitempoolactivitycreateAPIRequest {
	return &AlibabaretailmarketingitempoolactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivitycreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建商品池活动参数
func (r *AlibabaretailmarketingitempoolactivitycreateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitempoolactivitycreateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
