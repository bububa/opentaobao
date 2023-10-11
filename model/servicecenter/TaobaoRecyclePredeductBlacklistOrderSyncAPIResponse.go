package servicecenter

import (
	"encoding/xml"

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

// TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel is 同步服务商黑名单 成功返回结果
type TaobaoRecyclePredeductBlacklistOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_prededuct_blacklist_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步黑名单结果
	Data *OfnRecyclerSyncBlackListDto `json:"data,omitempty" xml:"data,omitempty"`
}
