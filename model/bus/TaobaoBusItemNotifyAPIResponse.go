package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusItemNotifyAPIResponse 汽车票城际巴士车次变更通知飞猪接口 API返回值
// taobao.bus.item.notify
//
// 供应商线路信息变更（如价格，发车时间，新增班次）同步到飞猪销售端需要 10分钟-10个小时的时间(跟供应商线路数量,接口可支持的并发量有关)。
// 为了让供应商线路信息变更能够及时体现到飞猪销售端，供应商可以在修改完自身系统的线路信息后，主动调用该接口通知飞猪，飞猪会将该线路数据进行一次同步。
type TaobaoBusItemNotifyAPIResponse struct {
	model.CommonResponse
	TaobaoBusItemNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusItemNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusItemNotifyAPIResponseModel).Reset()
}

// TaobaoBusItemNotifyAPIResponseModel is 汽车票城际巴士车次变更通知飞猪接口 成功返回结果
type TaobaoBusItemNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_item_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusItemNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusItemNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusItemNotifyAPIResponse)
	},
}

// GetTaobaoBusItemNotifyAPIResponse 从 sync.Pool 获取 TaobaoBusItemNotifyAPIResponse
func GetTaobaoBusItemNotifyAPIResponse() *TaobaoBusItemNotifyAPIResponse {
	return poolTaobaoBusItemNotifyAPIResponse.Get().(*TaobaoBusItemNotifyAPIResponse)
}

// ReleaseTaobaoBusItemNotifyAPIResponse 将 TaobaoBusItemNotifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusItemNotifyAPIResponse(v *TaobaoBusItemNotifyAPIResponse) {
	v.Reset()
	poolTaobaoBusItemNotifyAPIResponse.Put(v)
}
