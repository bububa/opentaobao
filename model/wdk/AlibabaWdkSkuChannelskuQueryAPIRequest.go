package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuQueryAPIRequest 查询渠道商品 API请求
// alibaba.wdk.sku.channelsku.query
//
// 查询渠道商品
type AlibabaWdkSkuChannelskuQueryAPIRequest struct {
	model.Params
	// 查询渠道商品的入参
	_param *ChannelSkuQueryDo
}

// NewAlibabaWdkSkuChannelskuQueryRequest 初始化AlibabaWdkSkuChannelskuQueryAPIRequest对象
func NewAlibabaWdkSkuChannelskuQueryRequest() *AlibabaWdkSkuChannelskuQueryAPIRequest {
	return &AlibabaWdkSkuChannelskuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuChannelskuQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询渠道商品的入参
func (r *AlibabaWdkSkuChannelskuQueryAPIRequest) SetParam(_param *ChannelSkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuChannelskuQueryAPIRequest) GetParam() *ChannelSkuQueryDo {
	return r._param
}

var poolAlibabaWdkSkuChannelskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuChannelskuQueryRequest()
	},
}

// GetAlibabaWdkSkuChannelskuQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuChannelskuQueryAPIRequest
func GetAlibabaWdkSkuChannelskuQueryAPIRequest() *AlibabaWdkSkuChannelskuQueryAPIRequest {
	return poolAlibabaWdkSkuChannelskuQueryAPIRequest.Get().(*AlibabaWdkSkuChannelskuQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuChannelskuQueryAPIRequest 将 AlibabaWdkSkuChannelskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuChannelskuQueryAPIRequest(v *AlibabaWdkSkuChannelskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuChannelskuQueryAPIRequest.Put(v)
}
