package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组基础报表数据对象 API返回值 
taobao.simba.rpt.adgroupbase.get

推广组基础报表数据对象
*/
type TaobaoSimbaRptAdgroupbaseGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptAdgroupbaseGetResponse
}

// 推广组基础报表数据对象 成功返回结果
type TaobaoSimbaRptAdgroupbaseGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_adgroupbase_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 广告组基础数据对象
    RptAdgroupBaseList   string `json:"rpt_adgroup_base_list,omitempty" xml:"rpt_adgroup_base_list,omitempty"`
}
