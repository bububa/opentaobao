package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeAddAPIResponse 房型新增接口（ID重复变更新） API返回值
// taobao.xhotel.roomtype.add
//
// 房型添加或更新
type TaobaoXhotelRoomtypeAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomtypeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomtypeAddAPIResponseModel).Reset()
}

// TaobaoXhotelRoomtypeAddAPIResponseModel is 房型新增接口（ID重复变更新） 成功返回结果
type TaobaoXhotelRoomtypeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房型信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xroomtype = nil
}

var poolTaobaoXhotelRoomtypeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomtypeAddAPIResponse)
	},
}

// GetTaobaoXhotelRoomtypeAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomtypeAddAPIResponse
func GetTaobaoXhotelRoomtypeAddAPIResponse() *TaobaoXhotelRoomtypeAddAPIResponse {
	return poolTaobaoXhotelRoomtypeAddAPIResponse.Get().(*TaobaoXhotelRoomtypeAddAPIResponse)
}

// ReleaseTaobaoXhotelRoomtypeAddAPIResponse 将 TaobaoXhotelRoomtypeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomtypeAddAPIResponse(v *TaobaoXhotelRoomtypeAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomtypeAddAPIResponse.Put(v)
}
