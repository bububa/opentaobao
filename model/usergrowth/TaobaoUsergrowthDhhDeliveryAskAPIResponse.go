package usergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthDhhDeliveryAskAPIResponse 广告曝光前判定接口V2 API返回值
// taobao.usergrowth.dhh.delivery.ask
//
// 提供给媒体在曝光广告前调用
type TaobaoUsergrowthDhhDeliveryAskAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthDhhDeliveryAskAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsergrowthDhhDeliveryAskAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsergrowthDhhDeliveryAskAPIResponseModel).Reset()
}

// TaobaoUsergrowthDhhDeliveryAskAPIResponseModel is 广告曝光前判定接口V2 成功返回结果
type TaobaoUsergrowthDhhDeliveryAskAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_dhh_delivery_ask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 在大航海平台可投放的任务ID列表
	TaskIdList []string `json:"task_id_list,omitempty" xml:"task_id_list>string,omitempty"`
	// 在大航海平台推荐的任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 错误码， 0： 成功；1：限流；2：服务不可用
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// true: 目标用户；false: 非目标用户
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsergrowthDhhDeliveryAskAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaskIdList = m.TaskIdList[:0]
	m.TaskId = ""
	m.Errcode = 0
	m.Result = false
}

var poolTaobaoUsergrowthDhhDeliveryAskAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsergrowthDhhDeliveryAskAPIResponse)
	},
}

// GetTaobaoUsergrowthDhhDeliveryAskAPIResponse 从 sync.Pool 获取 TaobaoUsergrowthDhhDeliveryAskAPIResponse
func GetTaobaoUsergrowthDhhDeliveryAskAPIResponse() *TaobaoUsergrowthDhhDeliveryAskAPIResponse {
	return poolTaobaoUsergrowthDhhDeliveryAskAPIResponse.Get().(*TaobaoUsergrowthDhhDeliveryAskAPIResponse)
}

// ReleaseTaobaoUsergrowthDhhDeliveryAskAPIResponse 将 TaobaoUsergrowthDhhDeliveryAskAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsergrowthDhhDeliveryAskAPIResponse(v *TaobaoUsergrowthDhhDeliveryAskAPIResponse) {
	v.Reset()
	poolTaobaoUsergrowthDhhDeliveryAskAPIResponse.Put(v)
}
