package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格计划rateplan删除 APIResponse
taobao.xhotel.rateplan.delete

酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
*/
type TaobaoXhotelRateplanDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRateplanDeleteResponse
}

type TaobaoXhotelRateplanDeleteResponse struct {
    XMLName xml.Name `xml:"xhotel_rateplan_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelRateplanDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
