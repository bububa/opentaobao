package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIQuerydetailAPIResponse 查面单号状态v1.0 API返回值
// taobao.wlb.waybill.i.querydetail
//
// 查看面单号的当前状态，如签收、发货、失效等。
type TaobaoWlbWaybillIQuerydetailAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillIQuerydetailAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIQuerydetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillIQuerydetailAPIResponseModel).Reset()
}

// TaobaoWlbWaybillIQuerydetailAPIResponseModel is 查面单号状态v1.0 成功返回结果
type TaobaoWlbWaybillIQuerydetailAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_querydetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 面单查询错误编码
	ErrorCodes []string `json:"error_codes,omitempty" xml:"error_codes>string,omitempty"`
	// 不存在的面单号
	InexistentWaybillCodes []string `json:"inexistent_waybill_codes,omitempty" xml:"inexistent_waybill_codes>string,omitempty"`
	// 面单详情
	WaybillDetails []WaybillDetailQueryInfo `json:"waybill_details,omitempty" xml:"waybill_details>waybill_detail_query_info,omitempty"`
	// 查询是否成功
	QuerySuccess bool `json:"query_success,omitempty" xml:"query_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIQuerydetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCodes = m.ErrorCodes[:0]
	m.InexistentWaybillCodes = m.InexistentWaybillCodes[:0]
	m.WaybillDetails = m.WaybillDetails[:0]
	m.QuerySuccess = false
}

var poolTaobaoWlbWaybillIQuerydetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillIQuerydetailAPIResponse)
	},
}

// GetTaobaoWlbWaybillIQuerydetailAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillIQuerydetailAPIResponse
func GetTaobaoWlbWaybillIQuerydetailAPIResponse() *TaobaoWlbWaybillIQuerydetailAPIResponse {
	return poolTaobaoWlbWaybillIQuerydetailAPIResponse.Get().(*TaobaoWlbWaybillIQuerydetailAPIResponse)
}

// ReleaseTaobaoWlbWaybillIQuerydetailAPIResponse 将 TaobaoWlbWaybillIQuerydetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillIQuerydetailAPIResponse(v *TaobaoWlbWaybillIQuerydetailAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillIQuerydetailAPIResponse.Put(v)
}
