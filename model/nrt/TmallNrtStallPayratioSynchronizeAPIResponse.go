package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStallPayratioSynchronizeAPIResponse 同步摊位收银比例 API返回值
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
type TmallNrtStallPayratioSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtStallPayratioSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtStallPayratioSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtStallPayratioSynchronizeAPIResponseModel).Reset()
}

// TmallNrtStallPayratioSynchronizeAPIResponseModel is 同步摊位收银比例 成功返回结果
type TmallNrtStallPayratioSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_stall_payratio_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtStallPayratioSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolTmallNrtStallPayratioSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtStallPayratioSynchronizeAPIResponse)
	},
}

// GetTmallNrtStallPayratioSynchronizeAPIResponse 从 sync.Pool 获取 TmallNrtStallPayratioSynchronizeAPIResponse
func GetTmallNrtStallPayratioSynchronizeAPIResponse() *TmallNrtStallPayratioSynchronizeAPIResponse {
	return poolTmallNrtStallPayratioSynchronizeAPIResponse.Get().(*TmallNrtStallPayratioSynchronizeAPIResponse)
}

// ReleaseTmallNrtStallPayratioSynchronizeAPIResponse 将 TmallNrtStallPayratioSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtStallPayratioSynchronizeAPIResponse(v *TmallNrtStallPayratioSynchronizeAPIResponse) {
	v.Reset()
	poolTmallNrtStallPayratioSynchronizeAPIResponse.Put(v)
}
