package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse 创建异步下载任务 API返回值
// taobao.universalbp.report.async.create.download.task
//
// 入参报表查询信息，出参下载任务id
type TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponseModel is 创建异步下载任务 成功返回结果
type TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_async_create_download_task_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportAsyncCreateDownloadTaskTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse)
	},
}

// GetTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse
func GetTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse() *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse {
	return poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse.Get().(*TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse)
}

// ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse 将 TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse(v *TaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportAsyncCreateDownloadTaskAPIResponse.Put(v)
}
