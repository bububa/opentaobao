package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂价格删除 APIResponse
taobao.xhotel.multiplerate.delete

酒店产品库rate删除
*/
type TaobaoXhotelMultiplerateDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelMultiplerateDeleteResponse
}

type TaobaoXhotelMultiplerateDeleteResponse struct {
    XMLName xml.Name `xml:"xhotel_multiplerate_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoXhotelMultiplerateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
