package fans

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansCashpoolCreateAPIResponse 创建资金池 API返回值
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
type TmallFansCashpoolCreateAPIResponse struct {
	model.CommonResponse
	TmallFansCashpoolCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFansCashpoolCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFansCashpoolCreateAPIResponseModel).Reset()
}

// TmallFansCashpoolCreateAPIResponseModel is 创建资金池 成功返回结果
type TmallFansCashpoolCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fans_cashpool_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	FansResult *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFansCashpoolCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FansResult = nil
}

var poolTmallFansCashpoolCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFansCashpoolCreateAPIResponse)
	},
}

// GetTmallFansCashpoolCreateAPIResponse 从 sync.Pool 获取 TmallFansCashpoolCreateAPIResponse
func GetTmallFansCashpoolCreateAPIResponse() *TmallFansCashpoolCreateAPIResponse {
	return poolTmallFansCashpoolCreateAPIResponse.Get().(*TmallFansCashpoolCreateAPIResponse)
}

// ReleaseTmallFansCashpoolCreateAPIResponse 将 TmallFansCashpoolCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallFansCashpoolCreateAPIResponse(v *TmallFansCashpoolCreateAPIResponse) {
	v.Reset()
	poolTmallFansCashpoolCreateAPIResponse.Put(v)
}
