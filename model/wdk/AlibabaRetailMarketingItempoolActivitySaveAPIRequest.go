package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivitysaveAPIRequest 【同城零售】保存商品池活动 API请求
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
type AlibabaretailmarketingitempoolactivitysaveAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaretailmarketingitempoolactivitysaveRequest 初始化AlibabaretailmarketingitempoolactivitysaveAPIRequest对象
func NewAlibabaretailmarketingitempoolactivitysaveRequest() *AlibabaretailmarketingitempoolactivitysaveAPIRequest {
	return &AlibabaretailmarketingitempoolactivitysaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivitysaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivitysaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivitysaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新商品池活动参数
func (r *AlibabaretailmarketingitempoolactivitysaveAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitempoolactivitysaveAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
