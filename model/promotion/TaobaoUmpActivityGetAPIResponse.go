package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitygetAPIResponse 查询营销活动 API返回值
// taobao.ump.activity.get
//
// 查询营销活动
type TaobaoumpactivitygetAPIResponse struct {
	model.CommonResponse
	TaobaoumpactivitygetAPIResponseModel
}

// TaobaoumpactivitygetAPIResponseModel is 查询营销活动 成功返回结果
type TaobaoumpactivitygetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动的内容，可以通过ump sdk中的marketingTool来完成对该内容的处理
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
