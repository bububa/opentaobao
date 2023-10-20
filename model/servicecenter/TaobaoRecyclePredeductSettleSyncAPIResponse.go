package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductSettleSyncAPIResponse 同步回收单线下打款明细 API返回值
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
type TaobaoRecyclePredeductSettleSyncAPIResponse struct {
	model.CommonResponse
	TaobaoRecyclePredeductSettleSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductSettleSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecyclePredeductSettleSyncAPIResponseModel).Reset()
}

// TaobaoRecyclePredeductSettleSyncAPIResponseModel is 同步回收单线下打款明细 成功返回结果
type TaobaoRecyclePredeductSettleSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_prededuct_settle_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步线下打款明细结果
	Data *OfnRecycleOrderSyncOfflineSettleInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductSettleSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecyclePredeductSettleSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecyclePredeductSettleSyncAPIResponse)
	},
}

// GetTaobaoRecyclePredeductSettleSyncAPIResponse 从 sync.Pool 获取 TaobaoRecyclePredeductSettleSyncAPIResponse
func GetTaobaoRecyclePredeductSettleSyncAPIResponse() *TaobaoRecyclePredeductSettleSyncAPIResponse {
	return poolTaobaoRecyclePredeductSettleSyncAPIResponse.Get().(*TaobaoRecyclePredeductSettleSyncAPIResponse)
}

// ReleaseTaobaoRecyclePredeductSettleSyncAPIResponse 将 TaobaoRecyclePredeductSettleSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecyclePredeductSettleSyncAPIResponse(v *TaobaoRecyclePredeductSettleSyncAPIResponse) {
	v.Reset()
	poolTaobaoRecyclePredeductSettleSyncAPIResponse.Put(v)
}
