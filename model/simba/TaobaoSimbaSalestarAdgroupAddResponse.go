package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(新)创建一个推广组 APIResponse
taobao.simba.salestar.adgroup.add

创建一个推广组
*/
type TaobaoSimbaSalestarAdgroupAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSalestarAdgroupAddResponse `json:"simba_salestar_adgroup_add_response,omitempty"` 
    TaobaoSimbaSalestarAdgroupAddResponse
}

/* model for simplify = false
type TaobaoSimbaSalestarAdgroupAddResponse struct {

    // 新增加的推广组
    
    Adgroup  *struct {
        ADGroup  *ADGroup `json:"ad_group,omitempty"`
    } `json:"adgroup,omitempty"`
    

}
*/

type TaobaoSimbaSalestarAdgroupAddResponse struct {

    // 新增加的推广组
    Adgroup   *ADGroup `json:"adgroup,omitempty"`

}
