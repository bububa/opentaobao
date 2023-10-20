package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocReceiveAPIResponse 电子保单数据接口 API返回值
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
type TmallServicecenterSpserviceorderEpocReceiveAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderEpocReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderEpocReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterSpserviceorderEpocReceiveAPIResponseModel).Reset()
}

// TmallServicecenterSpserviceorderEpocReceiveAPIResponseModel is 电子保单数据接口 成功返回结果
type TmallServicecenterSpserviceorderEpocReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_epoc_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderEpocReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterSpserviceorderEpocReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterSpserviceorderEpocReceiveAPIResponse)
	},
}

// GetTmallServicecenterSpserviceorderEpocReceiveAPIResponse 从 sync.Pool 获取 TmallServicecenterSpserviceorderEpocReceiveAPIResponse
func GetTmallServicecenterSpserviceorderEpocReceiveAPIResponse() *TmallServicecenterSpserviceorderEpocReceiveAPIResponse {
	return poolTmallServicecenterSpserviceorderEpocReceiveAPIResponse.Get().(*TmallServicecenterSpserviceorderEpocReceiveAPIResponse)
}

// ReleaseTmallServicecenterSpserviceorderEpocReceiveAPIResponse 将 TmallServicecenterSpserviceorderEpocReceiveAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterSpserviceorderEpocReceiveAPIResponse(v *TmallServicecenterSpserviceorderEpocReceiveAPIResponse) {
	v.Reset()
	poolTmallServicecenterSpserviceorderEpocReceiveAPIResponse.Put(v)
}
