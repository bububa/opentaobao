package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTaskCreateAPIResponse 聚石塔短信任务创建接口 API返回值
// taobao.jst.sms.task.create
//
// 聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
type TaobaoJstSmsTaskCreateAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsTaskCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstSmsTaskCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstSmsTaskCreateAPIResponseModel).Reset()
}

// TaobaoJstSmsTaskCreateAPIResponseModel is 聚石塔短信任务创建接口 成功返回结果
type TaobaoJstSmsTaskCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_task_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *SmsResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstSmsTaskCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoJstSmsTaskCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstSmsTaskCreateAPIResponse)
	},
}

// GetTaobaoJstSmsTaskCreateAPIResponse 从 sync.Pool 获取 TaobaoJstSmsTaskCreateAPIResponse
func GetTaobaoJstSmsTaskCreateAPIResponse() *TaobaoJstSmsTaskCreateAPIResponse {
	return poolTaobaoJstSmsTaskCreateAPIResponse.Get().(*TaobaoJstSmsTaskCreateAPIResponse)
}

// ReleaseTaobaoJstSmsTaskCreateAPIResponse 将 TaobaoJstSmsTaskCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstSmsTaskCreateAPIResponse(v *TaobaoJstSmsTaskCreateAPIResponse) {
	v.Reset()
	poolTaobaoJstSmsTaskCreateAPIResponse.Put(v)
}
