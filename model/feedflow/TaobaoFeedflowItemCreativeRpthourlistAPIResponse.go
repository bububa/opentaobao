package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcreativerpthourlistAPIResponse 超级推荐【商品推广】创意分时报表查询 API返回值
// taobao.feedflow.item.creative.rpthourlist
//
// 创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
type TaobaofeedflowitemcreativerpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcreativerpthourlistAPIResponseModel
}

// TaobaofeedflowitemcreativerpthourlistAPIResponseModel is 超级推荐【商品推广】创意分时报表查询 成功返回结果
type TaobaofeedflowitemcreativerpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_creative_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaofeedflowitemcreativerpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
