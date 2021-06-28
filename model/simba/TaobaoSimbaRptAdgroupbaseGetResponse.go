package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 APIResponse
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptAdgroupbaseGetResponse
}

type TaobaoSimbaRptAdgroupbaseGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_adgroupbase_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 广告组基础数据对象
    
    RptAdgroupBaseList   string `json:"rpt_adgroup_base_list,omitempty" xml:"rpt_adgroup_base_list,omitempty"`

    
}
