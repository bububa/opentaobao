package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivityskuaddAPIRequest 商品池活动新增商品 API请求
// alibaba.retail.marketing.itempool.activity.sku.add
//
// 新增或更新商品池活动商品信息【同城零售】
type AlibabaretailmarketingitempoolactivityskuaddAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// NewAlibabaretailmarketingitempoolactivityskuaddRequest 初始化AlibabaretailmarketingitempoolactivityskuaddAPIRequest对象
func NewAlibabaretailmarketingitempoolactivityskuaddRequest() *AlibabaretailmarketingitempoolactivityskuaddAPIRequest {
	return &AlibabaretailmarketingitempoolactivityskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailmarketingitempoolactivityskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailmarketingitempoolactivityskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailmarketingitempoolactivityskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaretailmarketingitempoolactivityskuaddAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaretailmarketingitempoolactivityskuaddAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
	return r._param
}
