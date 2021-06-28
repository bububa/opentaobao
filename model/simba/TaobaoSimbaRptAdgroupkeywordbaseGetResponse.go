package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词基础报表数据查询(明细数据不分类型查询) APIResponse
taobao.simba.rpt.adgroupkeywordbase.get

推广组下的词基础报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordbaseGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rpt_adgroupkeywordbase_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 词基础数据返回结果
    
    RptAdgroupkeywordBaseList   string `json:"rpt_adgroupkeyword_base_list,omitempty" xml:"