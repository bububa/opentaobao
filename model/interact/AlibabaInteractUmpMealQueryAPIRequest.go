package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractumpmealqueryAPIRequest 淘宝卖家搭配套餐查询 API请求
// alibaba.interact.ump.meal.query
//
// 查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
type AlibabainteractumpmealqueryAPIRequest struct {
	model.Params
}

// NewAlibabainteractumpmealqueryRequest 初始化AlibabainteractumpmealqueryAPIRequest对象
func NewAlibabainteractumpmealqueryRequest() *AlibabainteractumpmealqueryAPIRequest {
	return &AlibabainteractumpmealqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractumpmealqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.ump.meal.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractumpmealqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractumpmealqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
