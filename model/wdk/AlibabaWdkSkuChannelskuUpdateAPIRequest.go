package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuUpdateAPIRequest 更新渠道商品 API请求
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
type AlibabaWdkSkuChannelskuUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []ChannelSkuDo
}

// NewAlibabaWdkSkuChannelskuUpdateRequest 初始化AlibabaWdkSkuChannelskuUpdateAPIRequest对象
func NewAlibabaWdkSkuChannelskuUpdateRequest() *AlibabaWdkSkuChannelskuUpdateAPIRequest {
	return &AlibabaWdkSkuChannelskuUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuChannelskuUpdateAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuChannelskuUpdateAPIRequest) SetParamList(_paramList []ChannelSkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuChannelskuUpdateAPIRequest) GetParamList() []ChannelSkuDo {
	return r._paramList
}

var poolAlibabaWdkSkuChannelskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuChannelskuUpdateRequest()
	},
}

// GetAlibabaWdkSkuChannelskuUpdateRequest 从 sync.Pool 获取 AlibabaWdkSkuChannelskuUpdateAPIRequest
func GetAlibabaWdkSkuChannelskuUpdateAPIRequest() *AlibabaWdkSkuChannelskuUpdateAPIRequest {
	return poolAlibabaWdkSkuChannelskuUpdateAPIRequest.Get().(*AlibabaWdkSkuChannelskuUpdateAPIRequest)
}

// ReleaseAlibabaWdkSkuChannelskuUpdateAPIRequest 将 AlibabaWdkSkuChannelskuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuChannelskuUpdateAPIRequest(v *AlibabaWdkSkuChannelskuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuChannelskuUpdateAPIRequest.Put(v)
}
