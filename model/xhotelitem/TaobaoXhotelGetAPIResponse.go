package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelGetAPIResponse
酒店查询接口 API返回值
taobao.xhotel.get

酒店查询接口 */
type TaobaoXhotelGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelGetAPIResponseModel
}

// TaobaoXhotelGetAPIResponseModel is 酒店查询接口 成功返回结果
type TaobaoXhotelGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询得到的hotel
	Xhotel *FirstResult `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}
