package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMerchantItemCreatedraftAPIRequest
新建商品草稿 API请求
alibaba.wdk.merchant.item.createdraft

新建商品草稿erp接口 */
type AlibabaWdkMerchantItemCreatedraftAPIRequest struct {
	model.Params
	// 商品信息json
	_params string
}

// New
