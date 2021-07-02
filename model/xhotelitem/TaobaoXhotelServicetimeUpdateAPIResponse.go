package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelServicetimeUpdateAPIResponse 飞猪酒店多维度服务时间维护接口 API返回值
// taobao.xhotel.servicetime.update
//
// 飞猪酒店多维度服务时间维护，支持卖家维度，supplier维度，酒店维度
type TaobaoXhotelServicetimeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelServicetimeUpdateAPIResponseModel
}

// TaobaoXhotelServicetimeUpdateAPIResponseModel is 飞猪酒店多维度服务时间维护接口 成功返回结果
type TaobaoXhotelServicetimeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_servicetime_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelServicetimeUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
