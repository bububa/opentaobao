package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStallSynchronizeAPIResponse 摊位信息同步 API返回值
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
type TmallNrtStallSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtStallSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtStallSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtStallSynchronizeAPIResponseModel).Reset()
}

// TmallNrtStallSynchronizeAPIResponseModel is 摊位信息同步 成功返回结果
type TmallNrtStallSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_stall_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtStallSynchronize *ResultDo `json:"tmall_nrt_stall_synchronize,omitempty" xml:"tmall_nrt_stall_synchronize,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtStallSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmallNrtStallSynchronize = nil
}

var poolTmallNrtStallSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtStallSynchronizeAPIResponse)
	},
}

// GetTmallNrtStallSynchronizeAPIResponse 从 sync.Pool 获取 TmallNrtStallSynchronizeAPIResponse
func GetTmallNrtStallSynchronizeAPIResponse() *TmallNrtStallSynchronizeAPIResponse {
	return poolTmallNrtStallSynchronizeAPIResponse.Get().(*TmallNrtStallSynchronizeAPIResponse)
}

// ReleaseTmallNrtStallSynchronizeAPIResponse 将 TmallNrtStallSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtStallSynchronizeAPIResponse(v *TmallNrtStallSynchronizeAPIResponse) {
	v.Reset()
	poolTmallNrtStallSynchronizeAPIResponse.Put(v)
}
