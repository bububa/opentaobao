package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductListAPIRequest 供应商渠道产品列表查询 API请求
// alibaba.ascp.channel.supplier.product.list
//
// 供应商查询渠道产品列表
type AlibabaAscpChannelSupplierProductListAPIRequest struct {
	model.Params
	// 请求参数
	_param *ProductListQueryRequestForSupplier
}

// NewAlibabaAscpChannelSupplierProductListRequest 初始化AlibabaAscpChannelSupplierProductListAPIRequest对象
func NewAlibabaAscpChannelSupplierProductListRequest() *AlibabaAscpChannelSupplierProductListAPIRequest {
	return &AlibabaAscpChannelSupplierProductListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelSupplierProductListAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabaAscpChannelSupplierProductListAPIRequest) SetParam(_param *ProductListQueryRequestForSupplier) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAscpChannelSupplierProductListAPIRequest) GetParam() *ProductListQueryRequestForSupplier {
	return r._param
}

var poolAlibabaAscpChannelSupplierProductListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelSupplierProductListRequest()
	},
}

// GetAlibabaAscpChannelSupplierProductListRequest 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductListAPIRequest
func GetAlibabaAscpChannelSupplierProductListAPIRequest() *AlibabaAscpChannelSupplierProductListAPIRequest {
	return poolAlibabaAscpChannelSupplierProductListAPIRequest.Get().(*AlibabaAscpChannelSupplierProductListAPIRequest)
}

// ReleaseAlibabaAscpChannelSupplierProductListAPIRequest 将 AlibabaAscpChannelSupplierProductListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductListAPIRequest(v *AlibabaAscpChannelSupplierProductListAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductListAPIRequest.Put(v)
}
