package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderUpdateConfirmcodeAPIResponse 推送及更新订单确认号 API返回值
// taobao.xhotel.order.update.confirmcode
//
// 商家拿到订单确认号后，异步推送给飞猪或更新给飞猪。订单确认号用于到店查无时的紧急查单依据。
type TaobaoXhotelOrderUpdateConfirmcodeAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderUpdateConfirmcodeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderUpdateConfirmcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderUpdateConfirmcodeAPIResponseModel).Reset()
}

// TaobaoXhotelOrderUpdateConfirmcodeAPIResponseModel is 推送及更新订单确认号 成功返回结果
type TaobaoXhotelOrderUpdateConfirmcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_update_confirmcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 操作结果,成功返回success
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 是否操作成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderUpdateConfirmcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = ""
	m.Error = false
}

var poolTaobaoXhotelOrderUpdateConfirmcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderUpdateConfirmcodeAPIResponse)
	},
}

// GetTaobaoXhotelOrderUpdateConfirmcodeAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderUpdateConfirmcodeAPIResponse
func GetTaobaoXhotelOrderUpdateConfirmcodeAPIResponse() *TaobaoXhotelOrderUpdateConfirmcodeAPIResponse {
	return poolTaobaoXhotelOrderUpdateConfirmcodeAPIResponse.Get().(*TaobaoXhotelOrderUpdateConfirmcodeAPIResponse)
}

// ReleaseTaobaoXhotelOrderUpdateConfirmcodeAPIResponse 将 TaobaoXhotelOrderUpdateConfirmcodeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderUpdateConfirmcodeAPIResponse(v *TaobaoXhotelOrderUpdateConfirmcodeAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderUpdateConfirmcodeAPIResponse.Put(v)
}
