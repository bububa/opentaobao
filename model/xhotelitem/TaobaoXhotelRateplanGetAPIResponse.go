package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan查询 API返回值 
taobao.xhotel.rateplan.get

酒店产品库rateplan查询
*/
type TaobaoXhotelRateplanGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanGetAPIResponseModel
}

// 价格计划rateplan查询 成功返回结果
type TaobaoXhotelRateplanGetAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_rateplan_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // rateplan
    Rateplan   *RatePlan `json:"rateplan,omitempty" xml:"rateplan,omitempty"`
}
