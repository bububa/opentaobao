package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpPresalespackageConsignAPIResponse 预售预包尾款推单发货 API返回值
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
type AlibabaAscpPresalespackageConsignAPIResponse struct {
	model.CommonResponse
	AlibabaAscpPresalespackageConsignAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpPresalespackageConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpPresalespackageConsignAPIResponseModel).Reset()
}

// AlibabaAscpPresalespackageConsignAPIResponseModel is 预售预包尾款推单发货 成功返回结果
type AlibabaAscpPresalespackageConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_presalespackage_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *ErpPresaleFinalPayResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpPresalespackageConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpPresalespackageConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpPresalespackageConsignAPIResponse)
	},
}

// GetAlibabaAscpPresalespackageConsignAPIResponse 从 sync.Pool 获取 AlibabaAscpPresalespackageConsignAPIResponse
func GetAlibabaAscpPresalespackageConsignAPIResponse() *AlibabaAscpPresalespackageConsignAPIResponse {
	return poolAlibabaAscpPresalespackageConsignAPIResponse.Get().(*AlibabaAscpPresalespackageConsignAPIResponse)
}

// ReleaseAlibabaAscpPresalespackageConsignAPIResponse 将 AlibabaAscpPresalespackageConsignAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpPresalespackageConsignAPIResponse(v *AlibabaAscpPresalespackageConsignAPIResponse) {
	v.Reset()
	poolAlibabaAscpPresalespackageConsignAPIResponse.Put(v)
}
