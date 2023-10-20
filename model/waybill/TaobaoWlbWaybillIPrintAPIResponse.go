package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIPrintAPIResponse 打印确认接口v1.0 API返回值
// taobao.wlb.waybill.i.print
//
// 打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
type TaobaoWlbWaybillIPrintAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillIPrintAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIPrintAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillIPrintAPIResponseModel).Reset()
}

// TaobaoWlbWaybillIPrintAPIResponseModel is 打印确认接口v1.0 成功返回结果
type TaobaoWlbWaybillIPrintAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_print_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 面单打印信息
	WaybillApplyPrintCheckInfos []WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_infos,omitempty" xml:"waybill_apply_print_check_infos>waybill_apply_print_check_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillIPrintAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WaybillApplyPrintCheckInfos = m.WaybillApplyPrintCheckInfos[:0]
}

var poolTaobaoWlbWaybillIPrintAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillIPrintAPIResponse)
	},
}

// GetTaobaoWlbWaybillIPrintAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillIPrintAPIResponse
func GetTaobaoWlbWaybillIPrintAPIResponse() *TaobaoWlbWaybillIPrintAPIResponse {
	return poolTaobaoWlbWaybillIPrintAPIResponse.Get().(*TaobaoWlbWaybillIPrintAPIResponse)
}

// ReleaseTaobaoWlbWaybillIPrintAPIResponse 将 TaobaoWlbWaybillIPrintAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillIPrintAPIResponse(v *TaobaoWlbWaybillIPrintAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillIPrintAPIResponse.Put(v)
}
