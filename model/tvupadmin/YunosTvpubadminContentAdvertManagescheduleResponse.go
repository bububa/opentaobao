package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控修改 APIResponse
yunos.tvpubadmin.content.advert.manageschedule

广告牌照管控修改
*/
type YunosTvpubadminContentAdvertManagescheduleAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAdvertManagescheduleResponse
}

type YunosTvpubadminContentAdvertManagescheduleResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_manageschedule_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 管理广告排期
    
    Object   int64 `json:"object,omitempty" xml:"object,omitempty"`

    
}
