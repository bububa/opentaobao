package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse UD效果外投收藏加购明细拉取 API返回值
// alibaba.taobao.ud.smart.order.collect.cart.detail.pull
//
// UD效果外投收藏加购明细拉取
type AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse struct {
	model.CommonResponse
	AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponseModel).Reset()
}

// AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponseModel is UD效果外投收藏加购明细拉取 成功返回结果
type AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_ud_smart_order_collect_cart_detail_pull_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 收藏加购明细数据
	Result []OrderDetailCursorTopDTO2 `json:"result,omitempty" xml:"result>order_detail_cursor_top_dto2,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 页码游标，下一次请求传入，实现增量查询
	CursorId string `json:"cursor_id,omitempty" xml:"cursor_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.ErrorCode = ""
	m.ErrorMsg = ""
	m.CursorId = ""
	m.Success = false
}

var poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse)
	},
}

// GetAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse 从 sync.Pool 获取 AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse
func GetAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse() *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse {
	return poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse.Get().(*AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse)
}

// ReleaseAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse 将 AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse(v *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse) {
	v.Reset()
	poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse.Put(v)
}
