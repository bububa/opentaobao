package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbaseinforoomgetAPIResponse 酒店房型与房价查询 API返回值
// taobao.xhotel.baseinfo.room.get
//
// 根据outHid/hid获取酒店的房型和价格信息
type TaobaoxhotelbaseinforoomgetAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelbaseinforoomgetAPIResponseModel
}

// TaobaoxhotelbaseinforoomgetAPIResponseModel is 酒店房型与房价查询 成功返回结果
type TaobaoxhotelbaseinforoomgetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_baseinfo_room_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoxhotelbaseinforoomgetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
