package lstlogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstShiporderQueryAPIResponse 零售通发货单查询 API返回值
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
type AlibabaLstShiporderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstShiporderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstShiporderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstShiporderQueryAPIResponseModel).Reset()
}

// AlibabaLstShiporderQueryAPIResponseModel is 零售通发货单查询 成功返回结果
type AlibabaLstShiporderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_shiporder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstShiporderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstShiporderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstShiporderQueryAPIResponse)
	},
}

// GetAlibabaLstShiporderQueryAPIResponse 从 sync.Pool 获取 AlibabaLstShiporderQueryAPIResponse
func GetAlibabaLstShiporderQueryAPIResponse() *AlibabaLstShiporderQueryAPIResponse {
	return poolAlibabaLstShiporderQueryAPIResponse.Get().(*AlibabaLstShiporderQueryAPIResponse)
}

// ReleaseAlibabaLstShiporderQueryAPIResponse 将 AlibabaLstShiporderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstShiporderQueryAPIResponse(v *AlibabaLstShiporderQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstShiporderQueryAPIResponse.Put(v)
}
