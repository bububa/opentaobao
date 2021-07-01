package media

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInteractiveListGetbyuserAPIResponse
用户获取视频互动列表 API返回值
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表 */
type TaobaoInteractiveListGetbyuserAPIResponse struct {
	model.CommonResponse
	TaobaoInteractiveListGetbyuserAPIResponseModel
}

// TaobaoInteractiveListGetbyuserAPIResponseModel is 用户获取视频互动列表 成功返回结果
type TaobaoInteractiveListGetbyuserAPIResponseModel struct {
	XMLName xml.Name `xml:"interactive_list_getbyuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoInteractiveListGetbyuserResult `json:"result,omitempty" xml:"result,omitempty"`
}
