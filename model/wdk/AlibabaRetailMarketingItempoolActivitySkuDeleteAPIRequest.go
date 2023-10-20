package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivityskudeleteAPIRequest 删除商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
type AlibabaretailmarketingitempoolactivityskudeleteAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// NewAlibabaretailmarketingitempoolactivityskudeleteRequest 初始化AlibabaretailmarketingitempoolactivityskudeleteAPIRequest对象
func NewAlibabaretailmarketingitempoolactivityskudeleteRequest() *AlibabaretailmarketingitempoolactivityskudeleteAPIRequest {
	return &AlibabaretailmarketingitempoolactivityskudeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivityskudeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivityskudeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivityskudeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaretailmarketingitempoolactivityskudeleteAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitempoolactivityskudeleteAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
	return r._param
}
