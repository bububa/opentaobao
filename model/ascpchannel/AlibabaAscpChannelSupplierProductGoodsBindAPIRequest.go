package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductGoodsBindAPIRequest 渠道产品与货品绑定接口 API请求
// alibaba.ascp.channel.supplier.product.goods.bind
//
// 渠道产品与货品绑定接口
type AlibabaAscpChannelSupplierProductGoodsBindAPIRequest struct {
	model.Params
	// 请求
	_topBindProductGoodsRequest *TopBindProductGoodsRequest
}

// NewAlibabaAscpChannelSupplierProductGoodsBindRequest 初始化AlibabaAscpChannelSupplierProductGoodsBindAPIRequest对象
func NewAlibabaAscpChannelSupplierProductGoodsBindRequest() *AlibabaAscpChannelSupplierProductGoodsBindAPIRequest {
	return &AlibabaAscpChannelSupplierProductGoodsBindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) Reset() {
	r._topBindProductGoodsRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.goods.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopBindProductGoodsRequest is TopBindProductGoodsRequest Setter
// 请求
func (r *AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) SetTopBindProductGoodsRequest(_topBindProductGoodsRequest *TopBindProductGoodsRequest) error {
	r._topBindProductGoodsRequest = _topBindProductGoodsRequest
	r.Set("top_bind_product_goods_request", _topBindProductGoodsRequest)
	return nil
}

// GetTopBindProductGoodsRequest TopBindProductGoodsRequest Getter
func (r AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) GetTopBindProductGoodsRequest() *TopBindProductGoodsRequest {
	return r._topBindProductGoodsRequest
}

var poolAlibabaAscpChannelSupplierProductGoodsBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSupplierProductGoodsBindRequest()
	},
}

// GetAlibabaAscpChannelSupplierProductGoodsBindRequest 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductGoodsBindAPIRequest
func GetAlibabaAscpChannelSupplierProductGoodsBindAPIRequest() *AlibabaAscpChannelSupplierProductGoodsBindAPIRequest {
	return poolAlibabaAscpChannelSupplierProductGoodsBindAPIRequest.Get().(*AlibabaAscpChannelSupplierProductGoodsBindAPIRequest)
}

// ReleaseAlibabaAscpChannelSupplierProductGoodsBindAPIRequest 将 AlibabaAscpChannelSupplierProductGoodsBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductGoodsBindAPIRequest(v *AlibabaAscpChannelSupplierProductGoodsBindAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductGoodsBindAPIRequest.Put(v)
}
