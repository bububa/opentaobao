package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rate删除接口 APIResponse
taobao.xhotel.rate.delete

酒店产品库rate删除
*/
type TaobaoXhotelRateDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateDeleteResponse
}

type TaobaoXhotelRateDeleteResponse struct {
    XMLName xml.Name `xml:"xhotel_rate_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelRateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
