package wms

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsCainiaoBillQueryAPIResponse 查询单据列表 API返回值
// taobao.wlb.wms.cainiao.bill.query
//
// 查询单据列表
type TaobaoWlbWmsCainiaoBillQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWmsCainiaoBillQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWmsCainiaoBillQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWmsCainiaoBillQueryAPIResponseModel).Reset()
}

// TaobaoWlbWmsCainiaoBillQueryAPIResponseModel is 查询单据列表 成功返回结果
type TaobaoWlbWmsCainiaoBillQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_wms_cainiao_bill_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单列表信息
	OrderInfoList []CainiaoBillQueryOrderinfolist `json:"order_info_list,omitempty" xml:"order_info_list>cainiao_bill_query_orderinfolist,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWmsCainiaoBillQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderInfoList = m.OrderInfoList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbWmsCainiaoBillQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsCainiaoBillQueryAPIResponse)
	},
}

// GetTaobaoWlbWmsCainiaoBillQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbWmsCainiaoBillQueryAPIResponse
func GetTaobaoWlbWmsCainiaoBillQueryAPIResponse() *TaobaoWlbWmsCainiaoBillQueryAPIResponse {
	return poolTaobaoWlbWmsCainiaoBillQueryAPIResponse.Get().(*TaobaoWlbWmsCainiaoBillQueryAPIResponse)
}

// ReleaseTaobaoWlbWmsCainiaoBillQueryAPIResponse 将 TaobaoWlbWmsCainiaoBillQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWmsCainiaoBillQueryAPIResponse(v *TaobaoWlbWmsCainiaoBillQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbWmsCainiaoBillQueryAPIResponse.Put(v)
}
