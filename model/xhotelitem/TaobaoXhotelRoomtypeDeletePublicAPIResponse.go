package xhotelitem

import (
	"encoding/xml"

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

// TaobaoXhotelRoomtypeDeletePublicAPIResponseModel is 商家删除房型数据接口 成功返回结果
type TaobaoXhotelRoomtypeDeletePublicAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_delete_public_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelRoomtypeDeletePublicResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
