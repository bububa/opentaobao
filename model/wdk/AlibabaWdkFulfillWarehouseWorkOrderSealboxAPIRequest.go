package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest
仓封箱回告 API请求
alibaba.wdk.fulfill.warehouse.work.order.sealbox

仓封箱回告箱与包裹的关系 */
type AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest struct {
	model.Params
	// 同城交付物箱
	_sameTownBox *SameTownBox
}

// New
