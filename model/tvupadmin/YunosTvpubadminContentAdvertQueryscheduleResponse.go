package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控查询 APIResponse
yunos.tvpubadmin.content.advert.queryschedule

广告牌照管控查询
*/
type YunosTvpubadminContentAdvertQueryscheduleAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminContentAdvertQueryscheduleResponse
}

type YunosTvpubadminContentAdvertQueryscheduleResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_content_advert_queryschedule_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 具体数据信息
    
    Object   *PaginationDO `json:"object,omitempty" xml:"object,omitempty"`

    
}
