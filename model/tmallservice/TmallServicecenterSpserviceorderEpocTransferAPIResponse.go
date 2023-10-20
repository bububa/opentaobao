package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocTransferAPIResponse 电子保单受保人转移 API返回值
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
type TmallServicecenterSpserviceorderEpocTransferAPIResponse struct {
	model.CommonResponse
	TmallServicecenterSpserviceorderEpocTransferAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderEpocTransferAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterSpserviceorderEpocTransferAPIResponseModel).Reset()
}

// TmallServicecenterSpserviceorderEpocTransferAPIResponseModel is 电子保单受保人转移 成功返回结果
type TmallServicecenterSpserviceorderEpocTransferAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_spserviceorder_epoc_transfer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterSpserviceorderEpocTransferAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterSpserviceorderEpocTransferAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterSpserviceorderEpocTransferAPIResponse)
	},
}

// GetTmallServicecenterSpserviceorderEpocTransferAPIResponse 从 sync.Pool 获取 TmallServicecenterSpserviceorderEpocTransferAPIResponse
func GetTmallServicecenterSpserviceorderEpocTransferAPIResponse() *TmallServicecenterSpserviceorderEpocTransferAPIResponse {
	return poolTmallServicecenterSpserviceorderEpocTransferAPIResponse.Get().(*TmallServicecenterSpserviceorderEpocTransferAPIResponse)
}

// ReleaseTmallServicecenterSpserviceorderEpocTransferAPIResponse 将 TmallServicecenterSpserviceorderEpocTransferAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterSpserviceorderEpocTransferAPIResponse(v *TmallServicecenterSpserviceorderEpocTransferAPIResponse) {
	v.Reset()
	poolTmallServicecenterSpserviceorderEpocTransferAPIResponse.Put(v)
}
