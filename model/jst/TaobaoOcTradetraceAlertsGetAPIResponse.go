package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradetraceAlertsGetAPIResponse 异常订单信息获取 API返回值
// taobao.oc.tradetrace.alerts.get
//
// 提供订单预警模块的数据查询接口
type TaobaoOcTradetraceAlertsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOcTradetraceAlertsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOcTradetraceAlertsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOcTradetraceAlertsGetAPIResponseModel).Reset()
}

// TaobaoOcTradetraceAlertsGetAPIResponseModel is 异常订单信息获取 成功返回结果
type TaobaoOcTradetraceAlertsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_tradetrace_alerts_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常订单数据
	ResultList []SimpleAbnormalOrderDetail `json:"result_list,omitempty" xml:"result_list>simple_abnormal_order_detail,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOcTradetraceAlertsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
}

var poolTaobaoOcTradetraceAlertsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOcTradetraceAlertsGetAPIResponse)
	},
}

// GetTaobaoOcTradetraceAlertsGetAPIResponse 从 sync.Pool 获取 TaobaoOcTradetraceAlertsGetAPIResponse
func GetTaobaoOcTradetraceAlertsGetAPIResponse() *TaobaoOcTradetraceAlertsGetAPIResponse {
	return poolTaobaoOcTradetraceAlertsGetAPIResponse.Get().(*TaobaoOcTradetraceAlertsGetAPIResponse)
}

// ReleaseTaobaoOcTradetraceAlertsGetAPIResponse 将 TaobaoOcTradetraceAlertsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOcTradetraceAlertsGetAPIResponse(v *TaobaoOcTradetraceAlertsGetAPIResponse) {
	v.Reset()
	poolTaobaoOcTradetraceAlertsGetAPIResponse.Put(v)
}
