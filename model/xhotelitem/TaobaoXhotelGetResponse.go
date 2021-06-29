package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店查询接口 APIResponse
taobao.xhotel.get

酒店查询接口
*/
type TaobaoXhotelGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelGetResponse
}

type TaobaoXhotelGetResponse struct {
    XMLName xml.Name `xml:"xhotel_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询得到的hotel
    
    Xhotel   *FirstResult `json:"xhotel,omitempty" xml:"xhotel,omitempty"`

    
}
