package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广组 APIResponse
taobao.simba.adgroup.add

创建一个推广组
*/
type TaobaoSimbaAdgroupAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupAddResponse
}

type TaobaoSimbaAdgroupAddResponse struct {
    XMLName xml.Name `xml:"simba_adgroup_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 新增加的推广组
    
    Adgroup   *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`

    
}
