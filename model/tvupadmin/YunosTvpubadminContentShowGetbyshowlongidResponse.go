package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松根据节目longid获取节目元数据 APIResponse
yunos.tvpubadmin.content.show.getbyshowlongid

迎客松根据节目longid获取节目元数据
*/
type YunosTvpubadminContentShowGetbyshowlongidAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentShowGetbyshowlongidResponse
}

type YunosTvpubadminContentShowGetbyshowlongidResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getbyshowlongid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 节目元数据信息
    
    Object   string `json:"object,omitempty" xml:"object,omitempty"`

    
}
