package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCrowdRptdailylistAPIResponse
定向分日数据查询 API返回值
taobao.feedflow.item.crowd.rptdailylist

定向分日数据查询 */
type TaobaoFeedflowItemCrowdRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCrowdRptdailylistAPIResponseModel
}

// TaobaoFeedflowItemCrowdRptdailylistAPIResponseModel is 定向分日数据查询 成功返回结果
type TaobaoFeedflowItemCrowdRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowItemCrowdRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
