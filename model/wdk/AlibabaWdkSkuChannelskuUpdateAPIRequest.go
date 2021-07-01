package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuChannelskuUpdateAPIRequest
更新渠道商品 API请求
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入 */
type AlibabaWdkSkuChannelskuUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []ChannelSkuDo
}

// New
