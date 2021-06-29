package yunosappstore

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外投广告 APIResponse
yunos.appstore.open.getads

将广告外投给外部合作伙伴
*/
type YunosAppstoreOpenGetadsAPIResponse struct {
    model.CommonResponse
    YunosAppstoreOpenGetadsResponse
}

type YunosAppstoreOpenGetadsResponse struct {
    XMLName xml.Name `xml:"yunos_appstore_open_getads_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 请求id
    
    Rid   string `json:"rid,omitempty" xml:"rid,omitempty"`

    
    // 响应码
    
    Rc   int64 `json:"rc,omitempty" xml:"rc,omitempty"`

    
    // 响应消息
    
    Rm   string `json:"rm,omitempty" xml:"rm,omitempty"`

    
    // 广告集
    
    Ads   []AdInfo `json:"ads,omitempty" xml:"ads>ad_info,omitempty"`
    
    
}
