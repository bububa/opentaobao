package damai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletePaperformatAPIResponse 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API返回值
// alibaba.damai.mev.open.delete.paperformat
//
// deletePaperFormat
type AlibabaDamaiMevOpenDeletePaperformatAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletePaperformatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel).Reset()
}

// AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel is 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat 成功返回结果
type AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_delete_paperformat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletePaperformatResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMevOpenDeletePaperformatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMevOpenDeletePaperformatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMevOpenDeletePaperformatAPIResponse)
	},
}

// GetAlibabaDamaiMevOpenDeletePaperformatAPIResponse 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletePaperformatAPIResponse
func GetAlibabaDamaiMevOpenDeletePaperformatAPIResponse() *AlibabaDamaiMevOpenDeletePaperformatAPIResponse {
	return poolAlibabaDamaiMevOpenDeletePaperformatAPIResponse.Get().(*AlibabaDamaiMevOpenDeletePaperformatAPIResponse)
}

// ReleaseAlibabaDamaiMevOpenDeletePaperformatAPIResponse 将 AlibabaDamaiMevOpenDeletePaperformatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletePaperformatAPIResponse(v *AlibabaDamaiMevOpenDeletePaperformatAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletePaperformatAPIResponse.Put(v)
}
