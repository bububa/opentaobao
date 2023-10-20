package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuChannelskuAddAPIRequest 新增渠道商品 API请求
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
type AlibabaWdkSkuChannelskuAddAPIRequest struct {
	model.Params
	// 入参模型
	_chSkuDOList []ChannelSkuDo
}

// NewAlibabaWdkSkuChannelskuAddRequest 初始化AlibabaWdkSkuChannelskuAddAPIRequest对象
func NewAlibabaWdkSkuChannelskuAddRequest() *AlibabaWdkSkuChannelskuAddAPIRequest {
	return &AlibabaWdkSkuChannelskuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuChannelskuAddAPIRequest) Reset() {
	r._chSkuDOList = r._chSkuDOList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChSkuDOList is ChSkuDOList Setter
// 入参模型
func (r *AlibabaWdkSkuChannelskuAddAPIRequest) SetChSkuDOList(_chSkuDOList []ChannelSkuDo) error {
	r._chSkuDOList = _chSkuDOList
	r.Set("ch_sku_d_o_list", _chSkuDOList)
	return nil
}

// GetChSkuDOList ChSkuDOList Getter
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetChSkuDOList() []ChannelSkuDo {
	return r._chSkuDOList
}

var poolAlibabaWdkSkuChannelskuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuChannelskuAddRequest()
	},
}

// GetAlibabaWdkSkuChannelskuAddRequest 从 sync.Pool 获取 AlibabaWdkSkuChannelskuAddAPIRequest
func GetAlibabaWdkSkuChannelskuAddAPIRequest() *AlibabaWdkSkuChannelskuAddAPIRequest {
	return poolAlibabaWdkSkuChannelskuAddAPIRequest.Get().(*AlibabaWdkSkuChannelskuAddAPIRequest)
}

// ReleaseAlibabaWdkSkuChannelskuAddAPIRequest 将 AlibabaWdkSkuChannelskuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuChannelskuAddAPIRequest(v *AlibabaWdkSkuChannelskuAddAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuChannelskuAddAPIRequest.Put(v)
}
