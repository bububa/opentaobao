package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpPurchasePriceCreateAPIRequest
ascp采购价写入接口 API请求
alibaba.ascp.purchase.price.create

供应链平台采购价创建或修改接口 */
type AlibabaAscpPurchasePriceCreateAPIRequest struct {
	model.Params
	// 采购价创建/更新请求
	_createRequest *AlibabaAscpPurchasePriceCreateRequest
}

// New
