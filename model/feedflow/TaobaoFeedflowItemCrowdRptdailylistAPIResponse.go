package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcrowdrptdailylistAPIResponse 定向分日数据查询 API返回值
// taobao.feedflow.item.crowd.rptdailylist
//
// 定向分日数据查询
type TaobaofeedflowitemcrowdrptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcrowdrptdailylistAPIResponseModel
}

// TaobaofeedflowitemcrowdrptdailylistAPIResponseModel is 定向分日数据查询 成功返回结果
type TaobaofeedflowitemcrowdrptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_crowd_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaofeedflowitemcrowdrptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
