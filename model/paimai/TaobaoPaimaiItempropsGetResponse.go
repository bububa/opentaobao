package paimai

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
拍卖相关类目属性 APIResponse
taobao.paimai.itemprops.get

读取拍卖相关类目属性
*/
type TaobaoPaimaiItempropsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"paimai_itemprops_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
    
    ItemProps   []ItemProp `json:"item_props,omitempty" xml:"