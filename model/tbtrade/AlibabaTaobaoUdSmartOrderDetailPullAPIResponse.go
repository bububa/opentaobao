package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartOrderDetailPullAPIResponse UD效果外投订单明细拉取 API返回值
// alibaba.taobao.ud.smart.order.detail.pull
//
// UD效果外投订单明细拉取
type AlibabaTaobaoUdSmartOrderDetailPullAPIResponse struct {
	model.CommonResponse
	AlibabaTaobaoUdSmartOrderDetailPullAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartOrderDetailPullAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTaobaoUdSmartOrderDetailPullAPIResponseModel).Reset()
}

// AlibabaTaobaoUdSmartOrderDetailPullAPIResponseModel is UD效果外投订单明细拉取 成功返回结果
type AlibabaTaobaoUdSmartOrderDetailPullAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_ud_smart_order_detail_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单明细详情
	Result []OrderDetailCursorTopDto `json:"result,omitempty" xml:"result>order_detail_cursor_top_dto,omitempty"`
	//
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 页码游标，下一次请求传入，实现增量查询
	CursorId string `json:"cursor_id,omitempty" xml:"cursor_id,omitempty"`
	//
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	//
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartOrderDetailPullAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.ErrorCode = ""
	m.CursorId = ""
	m.ErrorMsg = ""
	m.Success = false
}

var poolAlibabaTaobaoUdSmartOrderDetailPullAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTaobaoUdSmartOrderDetailPullAPIResponse)
	},
}

// GetAlibabaTaobaoUdSmartOrderDetailPullAPIResponse 从 sync.Pool 获取 AlibabaTaobaoUdSmartOrderDetailPullAPIResponse
func GetAlibabaTaobaoUdSmartOrderDetailPullAPIResponse() *AlibabaTaobaoUdSmartOrderDetailPullAPIResponse {
	return poolAlibabaTaobaoUdSmartOrderDetailPullAPIResponse.Get().(*AlibabaTaobaoUdSmartOrderDetailPullAPIResponse)
}

// ReleaseAlibabaTaobaoUdSmartOrderDetailPullAPIResponse 将 AlibabaTaobaoUdSmartOrderDetailPullAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTaobaoUdSmartOrderDetailPullAPIResponse(v *AlibabaTaobaoUdSmartOrderDetailPullAPIResponse) {
	v.Reset()
	poolAlibabaTaobaoUdSmartOrderDetailPullAPIResponse.Put(v)
}
