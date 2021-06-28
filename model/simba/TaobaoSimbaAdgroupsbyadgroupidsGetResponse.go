package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量得到推广组 APIResponse
taobao.simba.adgroupsbyadgroupids.get

批量得到推广组
*/
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupsbyadgroupidsGetResponse
}

type TaobaoSimbaAdgroupsbyadgroupidsGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroupsbyadgroupids_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的推广组分页对象
    
    Adgroups   *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`

    
}
