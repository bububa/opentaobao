package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmctopicgroupdeleteAPIResponse 删除消息topic分组路由 API返回值
// taobao.tmc.topic.group.delete
//
// 删除根据topic名称路由消息到不同的分组关系
type TaobaotmctopicgroupdeleteAPIResponse struct {
	model.CommonResponse
	TaobaotmctopicgroupdeleteAPIResponseModel
}

// TaobaotmctopicgroupdeleteAPIResponseModel is 删除消息topic分组路由 成功返回结果
type TaobaotmctopicgroupdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_topic_group_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
