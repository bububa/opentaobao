package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbGetbycodeAPIResponse 根据营销积木块代码获取积木块 API返回值
// taobao.ump.mbb.getbycode
//
// 根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
type TaobaoUmpMbbGetbycodeAPIResponse struct {
	model.CommonResponse
	TaobaoUmpMbbGetbycodeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpMbbGetbycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpMbbGetbycodeAPIResponseModel).Reset()
}

// TaobaoUmpMbbGetbycodeAPIResponseModel is 根据营销积木块代码获取积木块 成功返回结果
type TaobaoUmpMbbGetbycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_mbb_getbycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销积木块的内容，通过ump sdk来进行处理
	Mbb string `json:"mbb,omitempty" xml:"mbb,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpMbbGetbycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Mbb = ""
}

var poolTaobaoUmpMbbGetbycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpMbbGetbycodeAPIResponse)
	},
}

// GetTaobaoUmpMbbGetbycodeAPIResponse 从 sync.Pool 获取 TaobaoUmpMbbGetbycodeAPIResponse
func GetTaobaoUmpMbbGetbycodeAPIResponse() *TaobaoUmpMbbGetbycodeAPIResponse {
	return poolTaobaoUmpMbbGetbycodeAPIResponse.Get().(*TaobaoUmpMbbGetbycodeAPIResponse)
}

// ReleaseTaobaoUmpMbbGetbycodeAPIResponse 将 TaobaoUmpMbbGetbycodeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpMbbGetbycodeAPIResponse(v *TaobaoUmpMbbGetbycodeAPIResponse) {
	v.Reset()
	poolTaobaoUmpMbbGetbycodeAPIResponse.Put(v)
}
