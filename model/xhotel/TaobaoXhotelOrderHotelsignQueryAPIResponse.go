package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderHotelsignQueryAPIResponse 获取直连酒店（客栈）签约上线进度信息 API返回值
// taobao.xhotel.order.hotelsign.query
//
// 获取直连酒店（客栈）签约上线进度信息
type TaobaoXhotelOrderHotelsignQueryAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderHotelsignQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderHotelsignQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderHotelsignQueryAPIResponseModel).Reset()
}

// TaobaoXhotelOrderHotelsignQueryAPIResponseModel is 获取直连酒店（客栈）签约上线进度信息 成功返回结果
type TaobaoXhotelOrderHotelsignQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_hotelsign_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// dsNhotelInfoDOList，当入参不包含hotelcode的时候，只有vendor的时候返回该对象
	DsNhotelInfoDOList []DsNhotelInfoDo `json:"ds_nhotel_info_d_o_list,omitempty" xml:"ds_nhotel_info_d_o_list>ds_nhotel_info_do,omitempty"`
	// outUuid
	OutUuid string `json:"out_uuid,omitempty" xml:"out_uuid,omitempty"`
	// hotelSignInfo，当入参中包含hotelcode和vendor的时候，返回该对象
	HotelSignInfo *DchotelSignDo `json:"hotel_sign_info,omitempty" xml:"hotel_sign_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderHotelsignQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DsNhotelInfoDOList = m.DsNhotelInfoDOList[:0]
	m.OutUuid = ""
	m.HotelSignInfo = nil
}

var poolTaobaoXhotelOrderHotelsignQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderHotelsignQueryAPIResponse)
	},
}

// GetTaobaoXhotelOrderHotelsignQueryAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderHotelsignQueryAPIResponse
func GetTaobaoXhotelOrderHotelsignQueryAPIResponse() *TaobaoXhotelOrderHotelsignQueryAPIResponse {
	return poolTaobaoXhotelOrderHotelsignQueryAPIResponse.Get().(*TaobaoXhotelOrderHotelsignQueryAPIResponse)
}

// ReleaseTaobaoXhotelOrderHotelsignQueryAPIResponse 将 TaobaoXhotelOrderHotelsignQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderHotelsignQueryAPIResponse(v *TaobaoXhotelOrderHotelsignQueryAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderHotelsignQueryAPIResponse.Put(v)
}
