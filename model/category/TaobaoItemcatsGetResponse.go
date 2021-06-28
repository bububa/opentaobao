package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取后台供卖家发布商品的标准商品类目 APIResponse
taobao.itemcats.get

获取后台供卖家发布商品的标准商品类目。
*/
type TaobaoItemcatsGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemcatsGetResponse
}

type TaobaoItemcatsGetResponse struct {
    XMLName xml.Name `xml:"itemcats_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最近修改时间(如果取增量，会返回该字段)。
    
    LastModified   string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`

    
    // 增量类目信息,根据fields传入的参数返回相应的结果；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
    
    ItemCats   []ItemCat `json:"item_cats,omitempty" xml:"item_cats>item_cat,omitempty"`
    
    
}
