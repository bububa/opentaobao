package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemStoreskustatusUpdateAPIRequest
修改门店商品状态 API请求
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态 */
type AlibabaWdkItemStoreskustatusUpdateAPIRequest struct {
	model.Params
	// bean
	_bean *UpdateStoreSkuLifeStatusRequestBean
}

// New
