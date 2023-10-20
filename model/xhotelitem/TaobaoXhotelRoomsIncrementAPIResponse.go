package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomsIncrementAPIResponse 房型库存推送接口（批量增量） API返回值
// taobao.xhotel.rooms.increment
//
// Room库存增量更新接口，用户仅需要更新ROOM库存中发生变化的库存日历即可。
type TaobaoXhotelRoomsIncrementAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomsIncrementAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomsIncrementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomsIncrementAPIResponseModel).Reset()
}

// TaobaoXhotelRoomsIncrementAPIResponseModel is 房型库存推送接口（批量增量） 成功返回结果
type TaobaoXhotelRoomsIncrementAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rooms_increment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功的gids LIST
	Gids []string `json:"gids,omitempty" xml:"gids>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomsIncrementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Gids = m.Gids[:0]
}

var poolTaobaoXhotelRoomsIncrementAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomsIncrementAPIResponse)
	},
}

// GetTaobaoXhotelRoomsIncrementAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomsIncrementAPIResponse
func GetTaobaoXhotelRoomsIncrementAPIResponse() *TaobaoXhotelRoomsIncrementAPIResponse {
	return poolTaobaoXhotelRoomsIncrementAPIResponse.Get().(*TaobaoXhotelRoomsIncrementAPIResponse)
}

// ReleaseTaobaoXhotelRoomsIncrementAPIResponse 将 TaobaoXhotelRoomsIncrementAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomsIncrementAPIResponse(v *TaobaoXhotelRoomsIncrementAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomsIncrementAPIResponse.Put(v)
}
