package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoGetAPIResponse 民宿查询营销活动 API返回值
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
type TaobaoXhotelBnbpromoGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoGetAPIResponseModel
}

// TaobaoXhotelBnbpromoGetAPIResponseModel is 民宿查询营销活动 成功返回结果
type TaobaoXhotelBnbpromoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbpromoGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
