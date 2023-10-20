package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelPmsGuestbillGetVtwoAPIResponse 客人PMS账单信息查询 API返回值
// taobao.xhotel.pms.guestbill.get.vtwo
//
// 从pms获取客人账单信息
type TaobaoXhotelPmsGuestbillGetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelPmsGuestbillGetVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel).Reset()
}

// TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel is 客人PMS账单信息查询 成功返回结果
type TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_pms_guestbill_get_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelPmsGuestbillGetVtwoResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelPmsGuestbillGetVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelPmsGuestbillGetVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelPmsGuestbillGetVtwoAPIResponse)
	},
}

// GetTaobaoXhotelPmsGuestbillGetVtwoAPIResponse 从 sync.Pool 获取 TaobaoXhotelPmsGuestbillGetVtwoAPIResponse
func GetTaobaoXhotelPmsGuestbillGetVtwoAPIResponse() *TaobaoXhotelPmsGuestbillGetVtwoAPIResponse {
	return poolTaobaoXhotelPmsGuestbillGetVtwoAPIResponse.Get().(*TaobaoXhotelPmsGuestbillGetVtwoAPIResponse)
}

// ReleaseTaobaoXhotelPmsGuestbillGetVtwoAPIResponse 将 TaobaoXhotelPmsGuestbillGetVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelPmsGuestbillGetVtwoAPIResponse(v *TaobaoXhotelPmsGuestbillGetVtwoAPIResponse) {
	v.Reset()
	poolTaobaoXhotelPmsGuestbillGetVtwoAPIResponse.Put(v)
}
