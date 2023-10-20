package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypeconflictdataAPIResponse 商家床型冲突数据接口 API返回值
// taobao.xhotel.roomtype.conflict.data
//
// 商家床型冲突数据接口
type TaobaoxhotelroomtypeconflictdataAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelroomtypeconflictdataAPIResponseModel
}

// TaobaoxhotelroomtypeconflictdataAPIResponseModel is 商家床型冲突数据接口 成功返回结果
type TaobaoxhotelroomtypeconflictdataAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_roomtype_conflict_data_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoxhotelroomtypeconflictdataResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
