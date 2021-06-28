package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词基础报表数据查询(明细数据不分类型查询) APIResponse
taobao.simba.rpt.adgroupkeywordbase.get

推广组下的词基础报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptAdgroupkeywordbaseGetResponse `json:"simba_rpt_adgroupkeywordbase_get_response,omitempty"` 
    TaobaoSimbaRptAdgroupkeywordbaseGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptAdgroupkeywordbaseGetResponse struct {

    // 词基础数据返回结果
    
    RptAdgroupkeywordBaseList   string `json:"rpt_adgroupkeyword_base_list,omitempty"`
    

}
*/

type TaobaoSimbaRptAdgroupkeywordbaseGetResponse struct {

    // 词基础数据返回结果
    RptAdgroupkeywordBaseList   string `json:"rpt_adgroupkeyword_base_list,omitempty"`

}
