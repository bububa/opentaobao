package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品是否推广 APIResponse
taobao.simba.adgroups.item.exist

判断在一个推广计划中是否已经推广了一个商品
*/
type TaobaoSimbaAdgroupsItemExistAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupsItemExistResponse
}

type TaobaoSimbaAdgroupsItemExistResponse struct {
    XMLName xml.Name `xml:"simba_adgroups_item_exist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true表示已经被推广，false表示没有被推广
    
    Exist   bool `json:"exist,omitempty" xml:"exist,omitempty"`

    
}
