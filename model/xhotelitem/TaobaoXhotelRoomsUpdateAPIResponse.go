package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomsUpdateAPIResponse 房型共享库存推送接口（批量全量） API返回值
// taobao.xhotel.rooms.update
//
// 此接口用于更新多个集市酒店商品房态信息，根据传入的gids更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。是全量更新，非增量，会把之前的房态进行覆盖。
type TaobaoXhotelRoomsUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomsUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomsUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomsUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRoomsUpdateAPIResponseModel is 房型共享库存推送接口（批量全量） 成功返回结果
type TaobaoXhotelRoomsUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rooms_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功的gids LIST
	Gids []string `json:"gids,omitempty" xml:"gids>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomsUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Gids = m.Gids[:0]
}

var poolTaobaoXhotelRoomsUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomsUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRoomsUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomsUpdateAPIResponse
func GetTaobaoXhotelRoomsUpdateAPIResponse() *TaobaoXhotelRoomsUpdateAPIResponse {
	return poolTaobaoXhotelRoomsUpdateAPIResponse.Get().(*TaobaoXhotelRoomsUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRoomsUpdateAPIResponse 将 TaobaoXhotelRoomsUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomsUpdateAPIResponse(v *TaobaoXhotelRoomsUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomsUpdateAPIResponse.Put(v)
}
