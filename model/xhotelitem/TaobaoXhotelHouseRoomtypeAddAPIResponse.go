package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelHouseRoomtypeAddAPIResponse 民宿房型信息添加 API返回值
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
type TaobaoXhotelHouseRoomtypeAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelHouseRoomtypeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelHouseRoomtypeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelHouseRoomtypeAddAPIResponseModel).Reset()
}

// TaobaoXhotelHouseRoomtypeAddAPIResponseModel is 民宿房型信息添加 成功返回结果
type TaobaoXhotelHouseRoomtypeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_house_roomtype_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房型信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelHouseRoomtypeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xroomtype = nil
}

var poolTaobaoXhotelHouseRoomtypeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelHouseRoomtypeAddAPIResponse)
	},
}

// GetTaobaoXhotelHouseRoomtypeAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelHouseRoomtypeAddAPIResponse
func GetTaobaoXhotelHouseRoomtypeAddAPIResponse() *TaobaoXhotelHouseRoomtypeAddAPIResponse {
	return poolTaobaoXhotelHouseRoomtypeAddAPIResponse.Get().(*TaobaoXhotelHouseRoomtypeAddAPIResponse)
}

// ReleaseTaobaoXhotelHouseRoomtypeAddAPIResponse 将 TaobaoXhotelHouseRoomtypeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelHouseRoomtypeAddAPIResponse(v *TaobaoXhotelHouseRoomtypeAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelHouseRoomtypeAddAPIResponse.Put(v)
}
