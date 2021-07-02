package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityCreateAPIRequest 创建商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
type AlibabaRetailMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivityCreateRequest 初始化AlibabaRetailMarketingItempoolActivityCreateAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityCreateRequest() *AlibabaRetailMarketingItempoolActivityCreateAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 创建商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivityCreateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
