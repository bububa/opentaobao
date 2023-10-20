package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstShiporderCancelAPIResponse 零售通发货单取消 API返回值
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
type AlibabaLstShiporderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaLstShiporderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstShiporderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstShiporderCancelAPIResponseModel).Reset()
}

// AlibabaLstShiporderCancelAPIResponseModel is 零售通发货单取消 成功返回结果
type AlibabaLstShiporderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_shiporder_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstShiporderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstShiporderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstShiporderCancelAPIResponse)
	},
}

// GetAlibabaLstShiporderCancelAPIResponse 从 sync.Pool 获取 AlibabaLstShiporderCancelAPIResponse
func GetAlibabaLstShiporderCancelAPIResponse() *AlibabaLstShiporderCancelAPIResponse {
	return poolAlibabaLstShiporderCancelAPIResponse.Get().(*AlibabaLstShiporderCancelAPIResponse)
}

// ReleaseAlibabaLstShiporderCancelAPIResponse 将 AlibabaLstShiporderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstShiporderCancelAPIResponse(v *AlibabaLstShiporderCancelAPIResponse) {
	v.Reset()
	poolAlibabaLstShiporderCancelAPIResponse.Put(v)
}
