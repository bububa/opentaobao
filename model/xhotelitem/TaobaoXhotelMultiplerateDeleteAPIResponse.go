package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateDeleteAPIResponse 复杂价格删除 API返回值
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
type TaobaoXhotelMultiplerateDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultiplerateDeleteAPIResponseModel
}

// TaobaoXhotelMultiplerateDeleteAPIResponseModel is 复杂价格删除 成功返回结果
type TaobaoXhotelMultiplerateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerate_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelMultiplerateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
