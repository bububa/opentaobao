package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyorderOrderCheckAPIResponse 自助改单服务发货订单校验 API返回值
// taobao.modifyorder.order.check
//
// 自助改单服务发货后订单回传接口
type TaobaoModifyorderOrderCheckAPIResponse struct {
	model.CommonResponse
	TaobaoModifyorderOrderCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoModifyorderOrderCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoModifyorderOrderCheckAPIResponseModel).Reset()
}

// TaobaoModifyorderOrderCheckAPIResponseModel is 自助改单服务发货订单校验 成功返回结果
type TaobaoModifyorderOrderCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"modifyorder_order_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoModifyorderOrderCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.Result = false
}

var poolTaobaoModifyorderOrderCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoModifyorderOrderCheckAPIResponse)
	},
}

// GetTaobaoModifyorderOrderCheckAPIResponse 从 sync.Pool 获取 TaobaoModifyorderOrderCheckAPIResponse
func GetTaobaoModifyorderOrderCheckAPIResponse() *TaobaoModifyorderOrderCheckAPIResponse {
	return poolTaobaoModifyorderOrderCheckAPIResponse.Get().(*TaobaoModifyorderOrderCheckAPIResponse)
}

// ReleaseTaobaoModifyorderOrderCheckAPIResponse 将 TaobaoModifyorderOrderCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoModifyorderOrderCheckAPIResponse(v *TaobaoModifyorderOrderCheckAPIResponse) {
	v.Reset()
	poolTaobaoModifyorderOrderCheckAPIResponse.Put(v)
}
