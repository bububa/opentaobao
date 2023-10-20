package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse 码上放心数据落地-获取每天日报 API返回值
// alibaba.alihealth.drug.download.getentdailytaskdtolist
//
// 码上放心数据落地-获取每天日报
type AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponseModel is 码上放心数据落地-获取每天日报 成功返回结果
type AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_download_getentdailytaskdtolist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的主要实体
	Model []EntDailyReportDto `json:"model,omitempty" xml:"model>ent_daily_report_dto,omitempty"`
	// Headers
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 请求回执信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求回执编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// http状态
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
	m.Headers = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.HttpStatusCode = 0
}

var poolAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse)
	},
}

// GetAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse
func GetAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse() *AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse {
	return poolAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse.Get().(*AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse)
}

// ReleaseAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse 将 AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse(v *AlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugDownloadGetentdailytaskdtolistAPIResponse.Put(v)
}
