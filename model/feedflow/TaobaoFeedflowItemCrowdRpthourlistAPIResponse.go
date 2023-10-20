package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdrpthourlistAPIResponse 超级推荐【商品推广】定向分时报表查询 API返回值
// taobao.feedflow.item.crowd.rpthourlist
//
// 广告主定向分时数据查询，支持广告主查询最近90天内某一天的定向维度分时报表数据
type TaobaofeedflowitemcrowdrpthourlistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcrowdrpthourlistAPIResponseModel
}

// TaobaofeedflowitemcrowdrpthourlistAPIResponseModel is 超级推荐【商品推广】定向分时报表查询 成功返回结果
type TaobaofeedflowitemcrowdrpthourlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_rpthourlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaofeedflowitemcrowdrpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
