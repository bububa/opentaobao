package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMsfReceiveAPIResponse 签收接口 API返回值
// tmall.msf.receive
//
// 签收接口
type TmallMsfReceiveAPIResponse struct {
	model.CommonResponse
	TmallMsfReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMsfReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMsfReceiveAPIResponseModel).Reset()
}

// TmallMsfReceiveAPIResponseModel is 签收接口 成功返回结果
type TmallMsfReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_msf_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMsfReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTmallMsfReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMsfReceiveAPIResponse)
	},
}

// GetTmallMsfReceiveAPIResponse 从 sync.Pool 获取 TmallMsfReceiveAPIResponse
func GetTmallMsfReceiveAPIResponse() *TmallMsfReceiveAPIResponse {
	return poolTmallMsfReceiveAPIResponse.Get().(*TmallMsfReceiveAPIResponse)
}

// ReleaseTmallMsfReceiveAPIResponse 将 TmallMsfReceiveAPIResponse 保存到 sync.Pool
func ReleaseTmallMsfReceiveAPIResponse(v *TmallMsfReceiveAPIResponse) {
	v.Reset()
	poolTmallMsfReceiveAPIResponse.Put(v)
}
