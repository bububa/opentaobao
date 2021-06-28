package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组下创意报表基础数据查询(汇总数据，不分类型) APIResponse
taobao.simba.rpt.adgroupcreativebase.get

推广组下创意报表基础数据查询(汇总数据，不分类型)
*/
type TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptAdgroupcreativebaseGetResponse `json:"simba_rpt_adgroupcreativebase_get_response,omitempty"` 
    TaobaoSimbaRptAdgroupcreativebaseGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptAdgroupcreativebaseGetResponse struct {

    // 推广组下的创意基础数据列表
    
    RptAdgroupcreativeBaseList   string `json:"rpt_adgroupcreative_base_list,omitempty"`
    

}
*/

type TaobaoSimbaRptAdgroupcreativebaseGetResponse struct {

    // 推广组下的创意基础数据列表
    RptAdgroupcreativeBaseList   string `json:"rpt_adgroupcreative_base_list,omitempty"`

}
