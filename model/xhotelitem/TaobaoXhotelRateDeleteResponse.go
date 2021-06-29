package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rate删除接口 API返回值 
taobao.xhotel.rate.delete

酒店产品库rate删除
*/
type TaobaoXhotelRateDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateDeleteResponse
}

// rate删除接口 成功返回结果
type TaobaoXhotelRateDeleteResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoXhotelRateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
