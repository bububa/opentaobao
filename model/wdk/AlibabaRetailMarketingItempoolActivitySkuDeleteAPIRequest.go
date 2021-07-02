package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest 删除商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest 初始化AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest() *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest {
	return &AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
	return r._param
}
