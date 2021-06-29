package baodian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务器时间获取 APIResponse
taobao.baodian.server.date.get

获取服务器时间
*/
type TaobaoBaodianServerDateGetAPIResponse struct {
    model.CommonResponse
    TaobaoBaodianServerDateGetResponse
}

type TaobaoBaodianServerDateGetResponse struct {
    XMLName xml.Name `xml:"baodian_server_date_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回时间为毫秒
    
    ServerDate   int64 `json:"server_date,omitempty" xml:"server_date,omitempty"`

    
}
