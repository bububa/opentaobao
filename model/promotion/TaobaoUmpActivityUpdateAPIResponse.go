package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivityupdateAPIResponse 修改活动信息 API返回值
// taobao.ump.activity.update
//
// 修改营销活动
type TaobaoumpactivityupdateAPIResponse struct {
	model.CommonResponse
	TaobaoumpactivityupdateAPIResponseModel
}

// TaobaoumpactivityupdateAPIResponseModel is 修改活动信息 成功返回结果
type TaobaoumpactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
