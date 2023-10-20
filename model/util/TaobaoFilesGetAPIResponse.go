package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilesGetAPIResponse 业务文件获取 API返回值
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
type TaobaoFilesGetAPIResponse struct {
	model.CommonResponse
	TaobaoFilesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilesGetAPIResponseModel).Reset()
}

// TaobaoFilesGetAPIResponseModel is 业务文件获取 成功返回结果
type TaobaoFilesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"files_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// results
	Results []TopDownloadRecordDo `json:"results,omitempty" xml:"results>top_download_record_do,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoFilesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilesGetAPIResponse)
	},
}

// GetTaobaoFilesGetAPIResponse 从 sync.Pool 获取 TaobaoFilesGetAPIResponse
func GetTaobaoFilesGetAPIResponse() *TaobaoFilesGetAPIResponse {
	return poolTaobaoFilesGetAPIResponse.Get().(*TaobaoFilesGetAPIResponse)
}

// ReleaseTaobaoFilesGetAPIResponse 将 TaobaoFilesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilesGetAPIResponse(v *TaobaoFilesGetAPIResponse) {
	v.Reset()
	poolTaobaoFilesGetAPIResponse.Put(v)
}
