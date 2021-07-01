package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBnbroomtypeAddAPIResponse
民宿新增房源 API返回值
taobao.xhotel.bnbroomtype.add

添加民宿房源 */
type TaobaoXhotelBnbroomtypeAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbroomtypeAddAPIResponseModel
}

// TaobaoXhotelBnbroomtypeAddAPIResponseModel is 民宿新增房源 成功返回结果
type TaobaoXhotelBnbroomtypeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbroomtype_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 房源信息
	Xroomtype *XRoomType `json:"xroomtype,omitempty" xml:"xroomtype,omitempty"`
}
