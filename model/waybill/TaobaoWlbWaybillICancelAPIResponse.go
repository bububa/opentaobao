package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillICancelAPIResponse 商家取消获取的电子面单号v1.0 API返回值
// taobao.wlb.waybill.i.cancel
//
// 面单号有误需要取消的时候，调用该接口取消获取的电子面单。
type TaobaoWlbWaybillICancelAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillICancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillICancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbWaybillICancelAPIResponseModel).Reset()
}

// TaobaoWlbWaybillICancelAPIResponseModel is 商家取消获取的电子面单号v1.0 成功返回结果
type TaobaoWlbWaybillICancelAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用取消是否成功
	CancelResult bool `json:"cancel_result,omitempty" xml:"cancel_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbWaybillICancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CancelResult = false
}

var poolTaobaoWlbWaybillICancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWaybillICancelAPIResponse)
	},
}

// GetTaobaoWlbWaybillICancelAPIResponse 从 sync.Pool 获取 TaobaoWlbWaybillICancelAPIResponse
func GetTaobaoWlbWaybillICancelAPIResponse() *TaobaoWlbWaybillICancelAPIResponse {
	return poolTaobaoWlbWaybillICancelAPIResponse.Get().(*TaobaoWlbWaybillICancelAPIResponse)
}

// ReleaseTaobaoWlbWaybillICancelAPIResponse 将 TaobaoWlbWaybillICancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbWaybillICancelAPIResponse(v *TaobaoWlbWaybillICancelAPIResponse) {
	v.Reset()
	poolTaobaoWlbWaybillICancelAPIResponse.Put(v)
}
