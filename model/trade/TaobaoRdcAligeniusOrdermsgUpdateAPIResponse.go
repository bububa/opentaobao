package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusOrdermsgUpdateAPIResponse 订单消息状态回传 API返回值
// taobao.rdc.aligenius.ordermsg.update
//
// 用于订单消息处理状态回传
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusOrdermsgUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel).Reset()
}

// TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel is 订单消息状态回传 成功返回结果
type TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"rdc_aligenius_ordermsg_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoRdcAligeniusOrdermsgUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRdcAligeniusOrdermsgUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoRdcAligeniusOrdermsgUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRdcAligeniusOrdermsgUpdateAPIResponse)
	},
}

// GetTaobaoRdcAligeniusOrdermsgUpdateAPIResponse 从 sync.Pool 获取 TaobaoRdcAligeniusOrdermsgUpdateAPIResponse
func GetTaobaoRdcAligeniusOrdermsgUpdateAPIResponse() *TaobaoRdcAligeniusOrdermsgUpdateAPIResponse {
	return poolTaobaoRdcAligeniusOrdermsgUpdateAPIResponse.Get().(*TaobaoRdcAligeniusOrdermsgUpdateAPIResponse)
}

// ReleaseTaobaoRdcAligeniusOrdermsgUpdateAPIResponse 将 TaobaoRdcAligeniusOrdermsgUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRdcAligeniusOrdermsgUpdateAPIResponse(v *TaobaoRdcAligeniusOrdermsgUpdateAPIResponse) {
	v.Reset()
	poolTaobaoRdcAligeniusOrdermsgUpdateAPIResponse.Put(v)
}
