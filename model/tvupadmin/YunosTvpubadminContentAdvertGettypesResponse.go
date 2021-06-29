package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取广告位类型 APIResponse
yunos.tvpubadmin.content.advert.gettypes

获取广告位类型
*/
type YunosTvpubadminContentAdvertGettypesAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAdvertGettypesResponse
}

type YunosTvpubadminContentAdvertGettypesResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_gettypes_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Map<Integer, String> json格式
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
