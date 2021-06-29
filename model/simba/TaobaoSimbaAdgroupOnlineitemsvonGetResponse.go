package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户上架在线销售的全部宝贝 APIResponse
taobao.simba.adgroup.onlineitemsvon.get

获取用户上架在线销售的全部宝贝
*/
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaAdgroupOnlineitemsvonGetResponse
}

type TaobaoSimbaAdgroupOnlineitemsvonGetResponse struct {
    XMLName xml.Name `xml:"simba_adgroup_onlineitemsvon_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 带分页的淘宝商品
    
    PageItem   *SubwayItemPartition `json:"page_item,omitempty" xml:"page_item,omitempty"`

    
}
