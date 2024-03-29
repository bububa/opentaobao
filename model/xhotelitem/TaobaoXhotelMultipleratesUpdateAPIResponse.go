package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultipleratesUpdateAPIResponse 复杂价格推送接口（批量全量） API返回值
// taobao.xhotel.multiplerates.update
//
// 批量更新复杂价格
// 涵盖了taobao.xhotel.rates.update的功能
type TaobaoXhotelMultipleratesUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultipleratesUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelMultipleratesUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelMultipleratesUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelMultipleratesUpdateAPIResponseModel is 复杂价格推送接口（批量全量） 成功返回结果
type TaobaoXhotelMultipleratesUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerates_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品id,房价id,入住人数，连住天数
	GidAndRpidOccupancyLengthofstay []string `json:"gid_and_rpid_occupancy_lengthofstay,omitempty" xml:"gid_and_rpid_occupancy_lengthofstay>string,omitempty"`
	// 批量更新的时候，如果部分更新失败，会展示部分失败的原因
	Warnmessage string `json:"warnmessage,omitempty" xml:"warnmessage,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelMultipleratesUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GidAndRpidOccupancyLengthofstay = m.GidAndRpidOccupancyLengthofstay[:0]
	m.Warnmessage = ""
}

var poolTaobaoXhotelMultipleratesUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelMultipleratesUpdateAPIResponse)
	},
}

// GetTaobaoXhotelMultipleratesUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelMultipleratesUpdateAPIResponse
func GetTaobaoXhotelMultipleratesUpdateAPIResponse() *TaobaoXhotelMultipleratesUpdateAPIResponse {
	return poolTaobaoXhotelMultipleratesUpdateAPIResponse.Get().(*TaobaoXhotelMultipleratesUpdateAPIResponse)
}

// ReleaseTaobaoXhotelMultipleratesUpdateAPIResponse 将 TaobaoXhotelMultipleratesUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelMultipleratesUpdateAPIResponse(v *TaobaoXhotelMultipleratesUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelMultipleratesUpdateAPIResponse.Put(v)
}
