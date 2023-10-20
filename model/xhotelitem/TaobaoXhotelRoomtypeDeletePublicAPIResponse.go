package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypedeletepublicAPIResponse 商家删除房型数据接口 API返回值
// taobao.xhotel.roomtype.delete.public
//
// 房型删除TOP接口
type TaobaoxhotelroomtypedeletepublicAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelroomtypedeletepublicAPIResponseModel
}

// TaobaoxhotelroomtypedeletepublicAPIResponseModel is 商家删除房型数据接口 成功返回结果
type TaobaoxhotelroomtypedeletepublicAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_delete_public_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoxhotelroomtypedeletepublicResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
