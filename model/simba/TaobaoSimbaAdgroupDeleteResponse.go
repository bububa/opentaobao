package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个推广组 APIResponse
taobao.simba.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupDeleteResponse
}

type TaobaoSimbaAdgroupDeleteResponse struct {
    XMLName xml.Name `xml:"simba_adgroup_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 被删除的推广组
    
    Adgroup   *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`

    
}
