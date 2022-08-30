package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySaveAPIRequest 【同城零售】保存商品池活动 API请求
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
type AlibabaRetailMarketingItempoolActivitySaveAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivitySaveRequest 初始化AlibabaRetailMarketingItempoolActivitySaveAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySaveRequest() *AlibabaRetailMarketingItempoolActivitySaveAPIRequest {
	return &AlibabaRetailMarketingItempoolActivitySaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 更新商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivitySaveAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}
