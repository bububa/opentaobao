package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApDeleteAPIResponse 删除价签AP设备 API返回值
// taobao.uscesl.biz.ap.delete
//
// 删除价签AP设备
type TaobaoUsceslBizApDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizApDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizApDeleteAPIResponseModel).Reset()
}

// TaobaoUsceslBizApDeleteAPIResponseModel is 删除价签AP设备 成功返回结果
type TaobaoUsceslBizApDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与否看result.success，返回true或者false
	Result *TaobaoUsceslBizApDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizApDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApDeleteAPIResponse)
	},
}

// GetTaobaoUsceslBizApDeleteAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizApDeleteAPIResponse
func GetTaobaoUsceslBizApDeleteAPIResponse() *TaobaoUsceslBizApDeleteAPIResponse {
	return poolTaobaoUsceslBizApDeleteAPIResponse.Get().(*TaobaoUsceslBizApDeleteAPIResponse)
}

// ReleaseTaobaoUsceslBizApDeleteAPIResponse 将 TaobaoUsceslBizApDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizApDeleteAPIResponse(v *TaobaoUsceslBizApDeleteAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizApDeleteAPIResponse.Put(v)
}
