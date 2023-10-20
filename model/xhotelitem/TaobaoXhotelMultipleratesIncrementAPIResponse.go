package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultipleratesIncrementAPIResponse 复杂房价推送接口（批量增量） API返回值
// taobao.xhotel.multiplerates.increment
//
// 复杂房价批量增量更新，只会更新指定日期的信息
// 完全涵盖了taobao.xhotel.rates.increment接口的功能
type TaobaoXhotelMultipleratesIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultipleratesIncrementAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelMultipleratesIncrementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelMultipleratesIncrementAPIResponseModel).Reset()
}

// TaobaoXhotelMultipleratesIncrementAPIResponseModel is 复杂房价推送接口（批量增量） 成功返回结果
type TaobaoXhotelMultipleratesIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerates_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品id-房价id-入住人数-连住天数  的集合
	GidAndRpidOccupancyLengthofstay []string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" xml:"gid_and_rpid_occupancy_lengthofstay>string,omitempty"`
	// 批量更新的时候，如果部分更新失败，会展示部分失败的原因
	Warnmessage string `json:"warnmessage,omitempty" xml:"warnmessage,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelMultipleratesIncrementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GidAndRpidOccupancyLengthofstay = m.GidAndRpidOccupancyLengthofstay[:0]
	m.Warnmessage = ""
}

var poolTaobaoXhotelMultipleratesIncrementAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelMultipleratesIncrementAPIResponse)
	},
}

// GetTaobaoXhotelMultipleratesIncrementAPIResponse 从 sync.Pool 获取 TaobaoXhotelMultipleratesIncrementAPIResponse
func GetTaobaoXhotelMultipleratesIncrementAPIResponse() *TaobaoXhotelMultipleratesIncrementAPIResponse {
	return poolTaobaoXhotelMultipleratesIncrementAPIResponse.Get().(*TaobaoXhotelMultipleratesIncrementAPIResponse)
}

// ReleaseTaobaoXhotelMultipleratesIncrementAPIResponse 将 TaobaoXhotelMultipleratesIncrementAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelMultipleratesIncrementAPIResponse(v *TaobaoXhotelMultipleratesIncrementAPIResponse) {
	v.Reset()
	poolTaobaoXhotelMultipleratesIncrementAPIResponse.Put(v)
}
