package idleitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderGetAPIResponse 闲鱼回收订单查询V2 API返回值
// alibaba.idle.recycle.order.get
//
// 闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
type AlibabaIdleRecycleOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleOrderGetAPIResponseModel).Reset()
}

// AlibabaIdleRecycleOrderGetAPIResponseModel is 闲鱼回收订单查询V2 成功返回结果
type AlibabaIdleRecycleOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败时候错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单信息
	Module *RecycleOrderTo `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Module = nil
	m.Result = false
}

var poolAlibabaIdleRecycleOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleOrderGetAPIResponse)
	},
}

// GetAlibabaIdleRecycleOrderGetAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleOrderGetAPIResponse
func GetAlibabaIdleRecycleOrderGetAPIResponse() *AlibabaIdleRecycleOrderGetAPIResponse {
	return poolAlibabaIdleRecycleOrderGetAPIResponse.Get().(*AlibabaIdleRecycleOrderGetAPIResponse)
}

// ReleaseAlibabaIdleRecycleOrderGetAPIResponse 将 AlibabaIdleRecycleOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleOrderGetAPIResponse(v *AlibabaIdleRecycleOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleOrderGetAPIResponse.Put(v)
}
