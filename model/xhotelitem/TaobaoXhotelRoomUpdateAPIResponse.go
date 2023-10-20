package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomUpdateAPIResponse 房型库存推送接口（全量推送） API返回值
// taobao.xhotel.room.update
//
// 此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
type TaobaoXhotelRoomUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRoomUpdateAPIResponseModel is 房型库存推送接口（全量推送） 成功返回结果
type TaobaoXhotelRoomUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_room_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// gid酒店商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Gid = 0
}

var poolTaobaoXhotelRoomUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRoomUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomUpdateAPIResponse
func GetTaobaoXhotelRoomUpdateAPIResponse() *TaobaoXhotelRoomUpdateAPIResponse {
	return poolTaobaoXhotelRoomUpdateAPIResponse.Get().(*TaobaoXhotelRoomUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRoomUpdateAPIResponse 将 TaobaoXhotelRoomUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomUpdateAPIResponse(v *TaobaoXhotelRoomUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomUpdateAPIResponse.Put(v)
}
