package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupdeleteAPIResponse 根据单元id删除单元 API返回值
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
type TaobaofeedflowitemadgroupdeleteAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemadgroupdeleteAPIResponseModel
}

// TaobaofeedflowitemadgroupdeleteAPIResponseModel is 根据单元id删除单元 成功返回结果
type TaobaofeedflowitemadgroupdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaofeedflowitemadgroupdeleteResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
