package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 APIResponse
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptAdgroupbaseGetResponse `json:"simba_rpt_adgroupbase_get_response,omitempty"` 
    TaobaoSimbaRptAdgroupbaseGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptAdgroupbaseGetResponse struct {

    // 广告组基础数据对象
    
    RptAdgroupBaseList   string `json:"rpt_adgroup_base_list,omitempty"`
    

}
*/

type TaobaoSimbaRptAdgroupbaseGetResponse struct {

    // 广告组基础数据对象
    RptAdgroupBaseList   string `json:"rpt_adgroup_base_list,omitempty"`

}
