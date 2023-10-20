package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFpcarRestpayReceiveAPIResponse 门店线下已收尾款 API返回值
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
type TmallCarFpcarRestpayReceiveAPIResponse struct {
	model.CommonResponse
	TmallCarFpcarRestpayReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarFpcarRestpayReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarFpcarRestpayReceiveAPIResponseModel).Reset()
}

// TmallCarFpcarRestpayReceiveAPIResponseModel is 门店线下已收尾款 成功返回结果
type TmallCarFpcarRestpayReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_fpcar_restpay_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 业务结果信息详细描述
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarFpcarRestpayReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Object = ""
	m.Succes = false
}

var poolTmallCarFpcarRestpayReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarFpcarRestpayReceiveAPIResponse)
	},
}

// GetTmallCarFpcarRestpayReceiveAPIResponse 从 sync.Pool 获取 TmallCarFpcarRestpayReceiveAPIResponse
func GetTmallCarFpcarRestpayReceiveAPIResponse() *TmallCarFpcarRestpayReceiveAPIResponse {
	return poolTmallCarFpcarRestpayReceiveAPIResponse.Get().(*TmallCarFpcarRestpayReceiveAPIResponse)
}

// ReleaseTmallCarFpcarRestpayReceiveAPIResponse 将 TmallCarFpcarRestpayReceiveAPIResponse 保存到 sync.Pool
func ReleaseTmallCarFpcarRestpayReceiveAPIResponse(v *TmallCarFpcarRestpayReceiveAPIResponse) {
	v.Reset()
	poolTmallCarFpcarRestpayReceiveAPIResponse.Put(v)
}
