package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuaddAPIRequest 新增渠道商品 API请求
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
type AlibabawdkskuchannelskuaddAPIRequest struct {
	model.Params
	// 入参模型
	_chSkuDOList []ChannelSkuDo
}

// NewAlibabawdkskuchannelskuaddRequest 初始化AlibabawdkskuchannelskuaddAPIRequest对象
func NewAlibabawdkskuchannelskuaddRequest() *AlibabawdkskuchannelskuaddAPIRequest {
	return &AlibabawdkskuchannelskuaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkskuchannelskuaddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.channelsku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkskuchannelskuaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkskuchannelskuaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChSkuDOList is ChSkuDOList Setter
// 入参模型
func (r *AlibabawdkskuchannelskuaddAPIRequest) SetChSkuDOList(_chSkuDOList []ChannelSkuDo) error {
	r._chSkuDOList = _chSkuDOList
	r.Set("ch_sku_d_o_list", _chSkuDOList)
	return nil
}

// GetChSkuDOList ChSkuDOList Getter
func (r AlibabawdkskuchannelskuaddAPIRequest) GetChSkuDOList() []ChannelSkuDo {
	return r._chSkuDOList
}
