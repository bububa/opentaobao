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

// New
