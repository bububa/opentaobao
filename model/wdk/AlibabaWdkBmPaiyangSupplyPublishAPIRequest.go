package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkBmPaiyangSupplyPublishAPIRequest
派样商品库存变更同步接口 API请求
alibaba.wdk.bm.paiyang.supply.publish

淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。 */
type AlibabaWdkBmPaiyangSupplyPublishAPIRequest struct {
	model.Params
	// 请求入参
	_isvSupplySyncParam *IsvSupplySyncParam
}

// New
