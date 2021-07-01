package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuChannelskuQueryAPIRequest
查询渠道商品 API请求
alibaba.wdk.sku.channelsku.query

查询渠道商品 */
type AlibabaWdkSkuChannelskuQueryAPIRequest struct {
	model.Params
	// 查询渠道商品的入参
	_param *ChannelSkuQueryDo
}

// New
