package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscToolCheckAPIResponse UMP工具检测 API返回值
// taobao.promotionmisc.tool.check
//
// UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
type TaobaoPromotionmiscToolCheckAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscToolCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscToolCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPromotionmiscToolCheckAPIResponseModel).Reset()
}

// TaobaoPromotionmiscToolCheckAPIResponseModel is UMP工具检测 成功返回结果
type TaobaoPromotionmiscToolCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_tool_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用错误信息描述。
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 工具审核结果。
	CheckToolModule *CheckToolModule `json:"check_tool_module,omitempty" xml:"check_tool_module,omitempty"`
	// 工具检测动作是否成功。检测是否通过请查看返回值default_model里的is_pass。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPromotionmiscToolCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMessage = ""
	m.CheckToolModule = nil
	m.IsSuccess = false
}

var poolTaobaoPromotionmiscToolCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPromotionmiscToolCheckAPIResponse)
	},
}

// GetTaobaoPromotionmiscToolCheckAPIResponse 从 sync.Pool 获取 TaobaoPromotionmiscToolCheckAPIResponse
func GetTaobaoPromotionmiscToolCheckAPIResponse() *TaobaoPromotionmiscToolCheckAPIResponse {
	return poolTaobaoPromotionmiscToolCheckAPIResponse.Get().(*TaobaoPromotionmiscToolCheckAPIResponse)
}

// ReleaseTaobaoPromotionmiscToolCheckAPIResponse 将 TaobaoPromotionmiscToolCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPromotionmiscToolCheckAPIResponse(v *TaobaoPromotionmiscToolCheckAPIResponse) {
	v.Reset()
	poolTaobaoPromotionmiscToolCheckAPIResponse.Put(v)
}
