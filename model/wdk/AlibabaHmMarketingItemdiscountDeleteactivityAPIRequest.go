package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitemdiscountdeleteactivityAPIRequest 删除商品特价活动 API请求
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabahmmarketingitemdiscountdeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityRequest
}

// NewAlibabahmmarketingitemdiscountdeleteactivityRequest 初始化AlibabahmmarketingitemdiscountdeleteactivityAPIRequest对象
func NewAlibabahmmarketingitemdiscountdeleteactivityRequest() *AlibabahmmarketingitemdiscountdeleteactivityAPIRequest {
	return &AlibabahmmarketingitemdiscountdeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitemdiscountdeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitemdiscountdeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitemdiscountdeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabahmmarketingitemdiscountdeleteactivityAPIRequest) SetParam(_param *CommonActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitemdiscountdeleteactivityAPIRequest) GetParam() *CommonActivityRequest {
	return r._param
}
