package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslBindAPIResponse 电子价签绑定接口 API返回值
// taobao.uscesl.biz.esl.bind
//
// 电子价签商品绑定接口
type TaobaoUsceslBizEslBindAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizEslBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizEslBindAPIResponseModel).Reset()
}

// TaobaoUsceslBizEslBindAPIResponseModel is 电子价签绑定接口 成功返回结果
type TaobaoUsceslBizEslBindAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_esl_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与否看result.success，返回true或者false
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslBizEslBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizEslBindAPIResponse)
	},
}

// GetTaobaoUsceslBizEslBindAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizEslBindAPIResponse
func GetTaobaoUsceslBizEslBindAPIResponse() *TaobaoUsceslBizEslBindAPIResponse {
	return poolTaobaoUsceslBizEslBindAPIResponse.Get().(*TaobaoUsceslBizEslBindAPIResponse)
}

// ReleaseTaobaoUsceslBizEslBindAPIResponse 将 TaobaoUsceslBizEslBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizEslBindAPIResponse(v *TaobaoUsceslBizEslBindAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizEslBindAPIResponse.Put(v)
}
