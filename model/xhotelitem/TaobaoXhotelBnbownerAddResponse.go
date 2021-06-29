package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿房东信息添加 APIResponse
taobao.xhotel.bnbowner.add

添加和更新民宿房东的信息
*/
type TaobaoXhotelBnbownerAddAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelBnbownerAddResponse
}

type TaobaoXhotelBnbownerAddResponse struct {
    XMLName xml.Name `xml:"xhotel_bnbowner_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果集
    
    Result   *TaobaoXhotelBnbownerAddResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
