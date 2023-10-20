package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopfeedlistgetAPIResponse 获取对话流列表 API返回值
// taobao.ailab.aicloud.top.feedlist.get
//
// 获取指定应用的对话流信息
type TaobaoailabaicloudtopfeedlistgetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopfeedlistgetAPIResponseModel
}

// TaobaoailabaicloudtopfeedlistgetAPIResponseModel is 获取对话流列表 成功返回结果
type TaobaoailabaicloudtopfeedlistgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_feedlist_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
