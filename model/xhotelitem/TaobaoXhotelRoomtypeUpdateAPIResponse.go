package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeUpdateAPIResponse 房型更新接口（ID不存在自动新增） API返回值
// taobao.xhotel.roomtype.update
//
// 酒店房型更新或添加
type TaobaoXhotelRoomtypeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomtypeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomtypeUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRoomtypeUpdateAPIResponseModel is 房型更新接口（ID不存在自动新增） 成功返回结果
type TaobaoXhotelRoomtypeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房型信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xroomtype = nil
}

var poolTaobaoXhotelRoomtypeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomtypeUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRoomtypeUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomtypeUpdateAPIResponse
func GetTaobaoXhotelRoomtypeUpdateAPIResponse() *TaobaoXhotelRoomtypeUpdateAPIResponse {
	return poolTaobaoXhotelRoomtypeUpdateAPIResponse.Get().(*TaobaoXhotelRoomtypeUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRoomtypeUpdateAPIResponse 将 TaobaoXhotelRoomtypeUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomtypeUpdateAPIResponse(v *TaobaoXhotelRoomtypeUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomtypeUpdateAPIResponse.Put(v)
}
