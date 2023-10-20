package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeDeletePublicAPIResponse 商家删除房型数据接口 API返回值
// taobao.xhotel.roomtype.delete.public
//
// 房型删除TOP接口
type TaobaoXhotelRoomtypeDeletePublicAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomtypeDeletePublicAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeDeletePublicAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRoomtypeDeletePublicAPIResponseModel).Reset()
}

// TaobaoXhotelRoomtypeDeletePublicAPIResponseModel is 商家删除房型数据接口 成功返回结果
type TaobaoXhotelRoomtypeDeletePublicAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_delete_public_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelRoomtypeDeletePublicResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRoomtypeDeletePublicAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelRoomtypeDeletePublicAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRoomtypeDeletePublicAPIResponse)
	},
}

// GetTaobaoXhotelRoomtypeDeletePublicAPIResponse 从 sync.Pool 获取 TaobaoXhotelRoomtypeDeletePublicAPIResponse
func GetTaobaoXhotelRoomtypeDeletePublicAPIResponse() *TaobaoXhotelRoomtypeDeletePublicAPIResponse {
	return poolTaobaoXhotelRoomtypeDeletePublicAPIResponse.Get().(*TaobaoXhotelRoomtypeDeletePublicAPIResponse)
}

// ReleaseTaobaoXhotelRoomtypeDeletePublicAPIResponse 将 TaobaoXhotelRoomtypeDeletePublicAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRoomtypeDeletePublicAPIResponse(v *TaobaoXhotelRoomtypeDeletePublicAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRoomtypeDeletePublicAPIResponse.Put(v)
}
