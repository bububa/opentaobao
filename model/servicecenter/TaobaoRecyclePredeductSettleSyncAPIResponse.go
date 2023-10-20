package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorecyclepredeductsettlesyncAPIResponse 同步回收单线下打款明细 API返回值
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
type TaobaorecyclepredeductsettlesyncAPIResponse struct {
	model.CommonResponse
	TaobaorecyclepredeductsettlesyncAPIResponseModel
}

// TaobaorecyclepredeductsettlesyncAPIResponseModel is 同步回收单线下打款明细 成功返回结果
type TaobaorecyclepredeductsettlesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_prededuct_settle_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步线下打款明细结果
	Data *OfnRecycleOrderSyncOfflineSettleInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}
