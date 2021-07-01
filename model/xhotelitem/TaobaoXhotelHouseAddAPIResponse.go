package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelHouseAddAPIResponse
非标准民宿房源添加 API返回值
taobao.xhotel.house.add

添加酒店或更新酒店 */
type TaobaoXhotelHouseAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelHouseAddAPIResponseModel
}

// TaobaoXhotelHouseAddAPIResponseModel is 非标准民宿房源添加 成功返回结果
type TaobaoXhotelHouseAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_house_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店信息
	Xhotel *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}
