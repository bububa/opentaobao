package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店更新接口（ID不存在自动新增） APIResponse
taobao.xhotel.update

酒店更新接口
*/
type TaobaoXhotelUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelUpdateResponse
}

type TaobaoXhotelUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 酒店信息
    
    Xhotel   *XHotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`

    
}
