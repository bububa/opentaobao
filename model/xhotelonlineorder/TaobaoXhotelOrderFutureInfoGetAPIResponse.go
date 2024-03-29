package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderFutureInfoGetAPIResponse 获取(查询)订单变更信息 API返回值
// taobao.xhotel.order.future.info.get
//
// 支持操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
type TaobaoXhotelOrderFutureInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderFutureInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderFutureInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderFutureInfoGetAPIResponseModel).Reset()
}

// TaobaoXhotelOrderFutureInfoGetAPIResponseModel is 获取(查询)订单变更信息 成功返回结果
type TaobaoXhotelOrderFutureInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_future_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回列表
	Results []HotelOrderDirectInfo `json:"results,omitempty" xml:"results>hotel_order_direct_info,omitempty"`
	// 返回外部请求流水号
	OutUuid string `json:"out_uuid,omitempty" xml:"out_uuid,omitempty"`
	// 是否成功标记
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderFutureInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.OutUuid = ""
	m.Issuccess = false
}

var poolTaobaoXhotelOrderFutureInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderFutureInfoGetAPIResponse)
	},
}

// GetTaobaoXhotelOrderFutureInfoGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderFutureInfoGetAPIResponse
func GetTaobaoXhotelOrderFutureInfoGetAPIResponse() *TaobaoXhotelOrderFutureInfoGetAPIResponse {
	return poolTaobaoXhotelOrderFutureInfoGetAPIResponse.Get().(*TaobaoXhotelOrderFutureInfoGetAPIResponse)
}

// ReleaseTaobaoXhotelOrderFutureInfoGetAPIResponse 将 TaobaoXhotelOrderFutureInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderFutureInfoGetAPIResponse(v *TaobaoXhotelOrderFutureInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderFutureInfoGetAPIResponse.Put(v)
}
