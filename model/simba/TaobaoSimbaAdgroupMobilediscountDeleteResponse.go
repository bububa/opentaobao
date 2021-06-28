package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量删除adgroup的移动溢价 APIResponse
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupMobilediscountDeleteResponse `json:"simba_adgroup_mobilediscount_delete_response,omitempty"` 
    TaobaoSimbaAdgroupMobilediscountDeleteResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupMobilediscountDeleteResponse struct {

    // 返回成功个数
    
    Result   int64 `json:"result,omitempty"`
    

    // 返回信息
    
    Message   string `json:"message,omitempty"`
    

    // 返回码
    
    Key   string `json:"key,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupMobilediscountDeleteResponse struct {

    // 返回成功个数
    Result   int64 `json:"result,omitempty"`

    // 返回信息
    Message   string `json:"message,omitempty"`

    // 返回码
    Key   string `json:"key,omitempty"`

}
