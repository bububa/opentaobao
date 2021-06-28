package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取标准类目属性值 APIResponse
taobao.itempropvalues.get

获取标准类目属性值
*/
type TaobaoItempropvaluesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItempropvaluesGetResponse `json:"itempropvalues_get_response,omitempty"` 
    TaobaoItempropvaluesGetResponse
}

/* model for simplify = false
type TaobaoItempropvaluesGetResponse struct {

    // 最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    
    LastModified   string `json:"last_modified,omitempty"`
    

    // 属性值,根据fields传入的参数返回相应的结果
    
    PropValues  struct {
        PropValue  []PropValue `json:"prop_value,omitempty"`
    } `json:"prop_values,omitempty"`
    

}
*/

type TaobaoItempropvaluesGetResponse struct {

    // 最近修改时间。格式:yyyy-MM-dd HH:mm:ss
    LastModified   string `json:"last_modified,omitempty"`

    // 属性值,根据fields传入的参数返回相应的结果
    PropValues   []PropValue `json:"prop_values,omitempty"`

}
