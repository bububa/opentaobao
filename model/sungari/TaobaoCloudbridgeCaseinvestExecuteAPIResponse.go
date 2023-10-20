package sungari

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCloudbridgeCaseinvestExecuteAPIResponse 红盾云桥案件协查服务 API返回值
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
type TaobaoCloudbridgeCaseinvestExecuteAPIResponse struct {
	model.CommonResponse
	TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCloudbridgeCaseinvestExecuteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel).Reset()
}

// TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel is 红盾云桥案件协查服务 成功返回结果
type TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel struct {
	XMLName xml.Name `xml:"cloudbridge_caseinvest_execute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoCloudbridgeCaseinvestExecuteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCloudbridgeCaseinvestExecuteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCloudbridgeCaseinvestExecuteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCloudbridgeCaseinvestExecuteAPIResponse)
	},
}

// GetTaobaoCloudbridgeCaseinvestExecuteAPIResponse 从 sync.Pool 获取 TaobaoCloudbridgeCaseinvestExecuteAPIResponse
func GetTaobaoCloudbridgeCaseinvestExecuteAPIResponse() *TaobaoCloudbridgeCaseinvestExecuteAPIResponse {
	return poolTaobaoCloudbridgeCaseinvestExecuteAPIResponse.Get().(*TaobaoCloudbridgeCaseinvestExecuteAPIResponse)
}

// ReleaseTaobaoCloudbridgeCaseinvestExecuteAPIResponse 将 TaobaoCloudbridgeCaseinvestExecuteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCloudbridgeCaseinvestExecuteAPIResponse(v *TaobaoCloudbridgeCaseinvestExecuteAPIResponse) {
	v.Reset()
	poolTaobaoCloudbridgeCaseinvestExecuteAPIResponse.Put(v)
}
