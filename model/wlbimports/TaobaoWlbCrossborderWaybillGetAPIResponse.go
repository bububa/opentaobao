package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbCrossborderWaybillGetAPIResponse 集货商家pdf和云打印面单获取，pdf需要配置白名单 API返回值
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//
//	需求链接：https://aone.alibaba-inc.com/req/21210808
type TaobaoWlbCrossborderWaybillGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbCrossborderWaybillGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbCrossborderWaybillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbCrossborderWaybillGetAPIResponseModel).Reset()
}

// TaobaoWlbCrossborderWaybillGetAPIResponseModel is 集货商家pdf和云打印面单获取，pdf需要配置白名单 成功返回结果
type TaobaoWlbCrossborderWaybillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_crossborder_waybill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *TaobaoWlbCrossborderWaybillGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbCrossborderWaybillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbCrossborderWaybillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbCrossborderWaybillGetAPIResponse)
	},
}

// GetTaobaoWlbCrossborderWaybillGetAPIResponse 从 sync.Pool 获取 TaobaoWlbCrossborderWaybillGetAPIResponse
func GetTaobaoWlbCrossborderWaybillGetAPIResponse() *TaobaoWlbCrossborderWaybillGetAPIResponse {
	return poolTaobaoWlbCrossborderWaybillGetAPIResponse.Get().(*TaobaoWlbCrossborderWaybillGetAPIResponse)
}

// ReleaseTaobaoWlbCrossborderWaybillGetAPIResponse 将 TaobaoWlbCrossborderWaybillGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbCrossborderWaybillGetAPIResponse(v *TaobaoWlbCrossborderWaybillGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbCrossborderWaybillGetAPIResponse.Put(v)
}
