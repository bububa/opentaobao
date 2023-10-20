package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse 同步服务商黑名单 API返回值
// taobao.recycle.prededuct.blacklist.order.sync
//
// 同步服务商黑名单
type TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse struct {
	model.CommonResponse
	TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel).Reset()
}

// TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel is 同步服务商黑名单 成功返回结果
type TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_prededuct_blacklist_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步黑名单结果
	Data *OfnRecyclerSyncBlackListDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse)
	},
}

// GetTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse 从 sync.Pool 获取 TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse
func GetTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse() *TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse {
	return poolTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse.Get().(*TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse)
}

// ReleaseTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse 将 TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse(v *TaobaoRecyclePredeductBlacklistOrderSyncAPIResponse) {
	v.Reset()
	poolTaobaoRecyclePredeductBlacklistOrderSyncAPIResponse.Put(v)
}
