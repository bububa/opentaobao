package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
对推广组进行单独移动溢价 APIResponse
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaAdgroupMobilediscountUpdateResponse `json:"simba_adgroup_mobilediscount_update_response,omitempty"` 
    TaobaoSimbaAdgroupMobilediscountUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaAdgroupMobilediscountUpdateResponse struct {

    // 更新成功的个数
    
    Result   int64 `json:"result,omitempty"`
    

    // 返回信息
    
    Message   string `json:"message,omitempty"`
    

    // 错误码
    
    Key   string `json:"key,omitempty"`
    

}
*/

type TaobaoSimbaAdgroupMobilediscountUpdateResponse struct {

    // 更新成功的个数
    Result   int64 `json:"result,omitempty"`

    // 返回信息
    Message   string `json:"message,omitempty"`

    // 错误码
    Key   string `json:"key,omitempty"`

}
