package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelhouseroomtypeaddAPIResponse 民宿房型信息添加 API返回值
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
type TaobaoxhotelhouseroomtypeaddAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelhouseroomtypeaddAPIResponseModel
}

// TaobaoxhotelhouseroomtypeaddAPIResponseModel is 民宿房型信息添加 成功返回结果
type TaobaoxhotelhouseroomtypeaddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_house_roomtype_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房型信息
	Xroomtype *XroomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
