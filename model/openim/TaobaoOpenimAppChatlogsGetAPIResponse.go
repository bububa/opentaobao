package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimappchatlogsgetAPIResponse openim应用聊天记录查询 API返回值
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
type TaobaoopenimappchatlogsgetAPIResponse struct {
	model.CommonResponse
	TaobaoopenimappchatlogsgetAPIResponseModel
}

// TaobaoopenimappchatlogsgetAPIResponseModel is openim应用聊天记录查询 成功返回结果
type TaobaoopenimappchatlogsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_app_chatlogs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *EsMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}
