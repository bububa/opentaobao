package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomGetAPIResponse room实体查询 API返回值
// taobao.xhotel.room.get
//
// 此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
type TaobaoXhotelRoomGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomGetAPIResponseModel).Reset()
}

// TaobaoXhotelRoomGetAPIResponseModel is room实体查询 成功返回结果
type TaobaoXhotelRoomGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_room_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房间信息
	Room *FirstResult `json:"room,omitempty" xml:"room,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Room = nil
}

var poolTaobaoXhotelRoomGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomGetAPIResponse)
	},
}

// GetTaobaoXhotelRoomGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomGetAPIResponse
func GetTaobaoXhotelRoomGetAPIResponse() *TaobaoXhotelRoomGetAPIResponse {
	return poolTaobaoXhotelRoomGetAPIResponse.Get().(*TaobaoXhotelRoomGetAPIResponse)
}

// ReleaseTaobaoXhotelRoomGetAPIResponse 将 TaobaoXhotelRoomGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomGetAPIResponse(v *TaobaoXhotelRoomGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomGetAPIResponse.Put(v)
}
