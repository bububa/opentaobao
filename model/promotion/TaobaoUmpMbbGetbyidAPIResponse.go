package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbGetbyidAPIResponse 获取营销积木块 API返回值
// taobao.ump.mbb.getbyid
//
// 根据积木块id获取营销积木块。
type TaobaoUmpMbbGetbyidAPIResponse struct {
	model.CommonResponse
	TaobaoUmpMbbGetbyidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpMbbGetbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpMbbGetbyidAPIResponseModel).Reset()
}

// TaobaoUmpMbbGetbyidAPIResponseModel is 获取营销积木块 成功返回结果
type TaobaoUmpMbbGetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_mbb_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销积木块定义信息，可以通过ump sdk里面的MBB.fromJson来处理
	Mbb string `json:"mbb,omitempty" xml:"mbb,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpMbbGetbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Mbb = ""
}

var poolTaobaoUmpMbbGetbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpMbbGetbyidAPIResponse)
	},
}

// GetTaobaoUmpMbbGetbyidAPIResponse 从 sync.Pool 获取 TaobaoUmpMbbGetbyidAPIResponse
func GetTaobaoUmpMbbGetbyidAPIResponse() *TaobaoUmpMbbGetbyidAPIResponse {
	return poolTaobaoUmpMbbGetbyidAPIResponse.Get().(*TaobaoUmpMbbGetbyidAPIResponse)
}

// ReleaseTaobaoUmpMbbGetbyidAPIResponse 将 TaobaoUmpMbbGetbyidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpMbbGetbyidAPIResponse(v *TaobaoUmpMbbGetbyidAPIResponse) {
	v.Reset()
	poolTaobaoUmpMbbGetbyidAPIResponse.Put(v)
}
