package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecycleorderfulfillsyncAPIResponse 同步回收单最终履约方式 API返回值
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
type TaobaorecycleorderfulfillsyncAPIResponse struct {
	model.CommonResponse
	TaobaorecycleorderfulfillsyncAPIResponseModel
}

// TaobaorecycleorderfulfillsyncAPIResponseModel is 同步回收单最终履约方式 成功返回结果
type TaobaorecycleorderfulfillsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_order_fulfill_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步回收单最终履约方式结果
	Data *OfnRecycleOrderSyncFulfillTypeDto `json:"data,omitempty" xml:"data,omitempty"`
}
