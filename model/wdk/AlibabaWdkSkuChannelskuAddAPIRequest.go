package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuChannelskuAddAPIRequest
新增渠道商品 API请求
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品 */
type AlibabaWdkSkuChannelskuAddAPIRequest struct {
	model.Params
	// 入参模型
	_chSkuDOList []ChannelSkuDo
}

// NewAlibabaWdkSkuChannelskuAddRequest 初始化AlibabaWdkSkuChannelskuAddAPIRequest对象
func NewAlibabaWdkSkuChannelskuAddRequest() *AlibabaWdkSkuChannelskuAddAPIRequest {
	return &AlibabaWdkSkuChannelskuAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChSkuDOList Setter
// 入参模型
func (r *AlibabaWdkSkuChannelskuAddAPIRequest) SetChSkuDOList(_chSkuDOList []ChannelSkuDo) error {
	r._chSkuDOList = _chSkuDOList
	r.Set("ch_sku_d_o_list", _chSkuDOList)
	return nil
}

// Get ChSkuDOList Getter
func (r AlibabaWdkSkuChannelskuAddAPIRequest) GetChSkuDOList() []ChannelSkuDo {
	return r._chSkuDOList
}
