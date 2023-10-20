package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabashopcategoryallgetAPIRequest 全部店铺分类信息查询接口 API请求
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
type AlibabashopcategoryallgetAPIRequest struct {
	model.Params
}

// NewAlibabashopcategoryallgetRequest 初始化AlibabashopcategoryallgetAPIRequest对象
func NewAlibabashopcategoryallgetRequest() *AlibabashopcategoryallgetAPIRequest {
	return &AlibabashopcategoryallgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabashopcategoryallgetAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.category.all.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabashopcategoryallgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabashopcategoryallgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
