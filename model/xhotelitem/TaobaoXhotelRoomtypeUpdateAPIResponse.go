package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomtypeUpdateAPIResponse
房型更新接口（ID不存在自动新增） API返回值
taobao.xhotel.roomtype.update

酒店房型更新或添加 */
type TaobaoXhotelRoomtypeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRoomtypeUpdateAPIResponseModel
}

// TaobaoXhotelRoomtypeUpdateAPIResponseModel is 房型更新接口（ID不存在自动新增） 成功返回结果
type TaobaoXhotelRoomtypeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房型信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
