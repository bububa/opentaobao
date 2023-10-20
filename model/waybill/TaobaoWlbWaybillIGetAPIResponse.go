package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIGetAPIResponse 获取物流服务商电子面单号v1.0 API返回值
// taobao.wlb.waybill.i.get
//
// 商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
type TaobaoWlbWaybillIGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillIGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillIGetAPIResponseModel).Reset()
}

// TaobaoWlbWaybillIGetAPIResponseModel is 获取物流服务商电子面单号v1.0 成功返回结果
type TaobaoWlbWaybillIGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 面单申请接口返回信息
	WaybillApplyNewCols []WaybillApplyNewInfo `json:"waybill_apply_new_cols,omitempty" xml:"waybill_apply_new_cols>waybill_apply_new_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WaybillApplyNewCols = m.WaybillApplyNewCols[:0]
}

var poolTaobaoWlbWaybillIGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillIGetAPIResponse)
	},
}

// GetTaobaoWlbWaybillIGetAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillIGetAPIResponse
func GetTaobaoWlbWaybillIGetAPIResponse() *TaobaoWlbWaybillIGetAPIResponse {
	return poolTaobaoWlbWaybillIGetAPIResponse.Get().(*TaobaoWlbWaybillIGetAPIResponse)
}

// ReleaseTaobaoWlbWaybillIGetAPIResponse 将 TaobaoWlbWaybillIGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillIGetAPIResponse(v *TaobaoWlbWaybillIGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillIGetAPIResponse.Put(v)
}
