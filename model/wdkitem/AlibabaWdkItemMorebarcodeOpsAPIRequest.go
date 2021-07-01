package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMorebarcodeOpsAPIRequest
商品一品多码维护操作 API请求
alibaba.wdk.item.morebarcode.ops

商品一品多码维护操作 */
type AlibabaWdkItemMorebarcodeOpsAPIRequest struct {
	model.Params
	// bean
	_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean
}

// New
