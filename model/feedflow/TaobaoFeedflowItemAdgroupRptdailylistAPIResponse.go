package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgrouprptdailylistAPIResponse 推广单元分日数据查询 API返回值
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
type TaobaofeedflowitemadgrouprptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemadgrouprptdailylistAPIResponseModel
}

// TaobaofeedflowitemadgrouprptdailylistAPIResponseModel is 推广单元分日数据查询 成功返回结果
type TaobaofeedflowitemadgrouprptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaofeedflowitemadgrouprptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
