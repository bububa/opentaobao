package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarContractDownloadAPIResponse 合同下载 API返回值
// tmall.car.contract.download
//
// 目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
// 因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
type TmallCarContractDownloadAPIResponse struct {
	model.CommonResponse
	TmallCarContractDownloadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarContractDownloadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarContractDownloadAPIResponseModel).Reset()
}

// TmallCarContractDownloadAPIResponseModel is 合同下载 成功返回结果
type TmallCarContractDownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_contract_download_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TmallCarContractDownloadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarContractDownloadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarContractDownloadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarContractDownloadAPIResponse)
	},
}

// GetTmallCarContractDownloadAPIResponse 从 sync.Pool 获取 TmallCarContractDownloadAPIResponse
func GetTmallCarContractDownloadAPIResponse() *TmallCarContractDownloadAPIResponse {
	return poolTmallCarContractDownloadAPIResponse.Get().(*TmallCarContractDownloadAPIResponse)
}

// ReleaseTmallCarContractDownloadAPIResponse 将 TmallCarContractDownloadAPIResponse 保存到 sync.Pool
func ReleaseTmallCarContractDownloadAPIResponse(v *TmallCarContractDownloadAPIResponse) {
	v.Reset()
	poolTmallCarContractDownloadAPIResponse.Put(v)
}
