package tbrefund

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallDisputeReceiveGetAPIResponse 天猫逆向纠纷查询 API返回值
// tmall.dispute.receive.get
//
// 展示商家所有退款信息
type TmallDisputeReceiveGetAPIResponse struct {
	model.CommonResponse
	TmallDisputeReceiveGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallDisputeReceiveGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallDisputeReceiveGetAPIResponseModel).Reset()
}

// TmallDisputeReceiveGetAPIResponseModel is 天猫逆向纠纷查询 成功返回结果
type TmallDisputeReceiveGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_dispute_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallDisputeReceiveGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallDisputeReceiveGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallDisputeReceiveGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallDisputeReceiveGetAPIResponse)
	},
}

// GetTmallDisputeReceiveGetAPIResponse 从 sync.Pool 获取 TmallDisputeReceiveGetAPIResponse
func GetTmallDisputeReceiveGetAPIResponse() *TmallDisputeReceiveGetAPIResponse {
	return poolTmallDisputeReceiveGetAPIResponse.Get().(*TmallDisputeReceiveGetAPIResponse)
}

// ReleaseTmallDisputeReceiveGetAPIResponse 将 TmallDisputeReceiveGetAPIResponse 保存到 sync.Pool
func ReleaseTmallDisputeReceiveGetAPIResponse(v *TmallDisputeReceiveGetAPIResponse) {
	v.Reset()
	poolTmallDisputeReceiveGetAPIResponse.Put(v)
}
