package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest 获取商品类型配置项 API请求
// alibaba.onetouch.logistics.express.special.product.type.list
//
// 获取商品类型配置项
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest struct {
	model.Params
}

// NewAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest 初始化AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.special.product.type.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
