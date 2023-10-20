package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDaogoubaoOrderStatisticsTotalAPIResponse 销售订单总额统计 API返回值
// taobao.daogoubao.order.statistics.total
//
// 对接千牛端数字中心
type TaobaoDaogoubaoOrderStatisticsTotalAPIResponse struct {
	model.CommonResponse
	TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDaogoubaoOrderStatisticsTotalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel).Reset()
}

// TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel is 销售订单总额统计 成功返回结果
type TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel struct {
	XMLName xml.Name `xml:"daogoubao_order_statistics_total_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderStatisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDaogoubaoOrderStatisticsTotalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDaogoubaoOrderStatisticsTotalAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDaogoubaoOrderStatisticsTotalAPIResponse)
	},
}

// GetTaobaoDaogoubaoOrderStatisticsTotalAPIResponse 从 sync.Pool 获取 TaobaoDaogoubaoOrderStatisticsTotalAPIResponse
func GetTaobaoDaogoubaoOrderStatisticsTotalAPIResponse() *TaobaoDaogoubaoOrderStatisticsTotalAPIResponse {
	return poolTaobaoDaogoubaoOrderStatisticsTotalAPIResponse.Get().(*TaobaoDaogoubaoOrderStatisticsTotalAPIResponse)
}

// ReleaseTaobaoDaogoubaoOrderStatisticsTotalAPIResponse 将 TaobaoDaogoubaoOrderStatisticsTotalAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDaogoubaoOrderStatisticsTotalAPIResponse(v *TaobaoDaogoubaoOrderStatisticsTotalAPIResponse) {
	v.Reset()
	poolTaobaoDaogoubaoOrderStatisticsTotalAPIResponse.Put(v)
}
