package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除酒店接口 APIResponse
taobao.xhotel.delete

删除飞猪酒店数据接口
*/
type TaobaoXhotelDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelDeleteResponse
}

type TaobaoXhotelDeleteResponse struct {
    XMLName xml.Name `xml:"xhotel_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 删除结果
    
    Result   *TaobaoXhotelDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
