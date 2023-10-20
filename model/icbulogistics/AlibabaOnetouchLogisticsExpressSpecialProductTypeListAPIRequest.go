package icbulogistics

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.special.product.type.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest()
	},
}

// GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListRequest 从 sync.Pool 获取 AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest
func GetAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest() *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest {
	return poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest.Get().(*AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest)
}

// ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest 将 AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest 放入 sync.Pool
func ReleaseAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest(v *AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest) {
	v.Reset()
	poolAlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIRequest.Put(v)
}
