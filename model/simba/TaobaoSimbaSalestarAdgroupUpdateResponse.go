package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新一个推广组的信息 APIResponse
taobao.simba.salestar.adgroup.update

更新一个推广组的信息，可以设置 是否上线
*/
type TaobaoSimbaSalestarAdgroupUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarAdgroupUpdateResponse
}

type TaobaoSimbaSalestarAdgroupUpdateResponse struct {
    XMLName xml.Name `xml:"simba_salestar_adgroup_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 被修改的推广组
    
    Adgroup   *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`

    
}
