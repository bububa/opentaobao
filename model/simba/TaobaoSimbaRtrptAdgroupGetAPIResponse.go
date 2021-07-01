package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广组实时报表数据 API返回值 
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
type TaobaoSimbaRtrptAdgroupGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRtrptAdgroupGetAPIResponseModel
}

// 获取推广组实时报表数据 成功返回结果
type TaobaoSimbaRtrptAdgroupGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_rtrpt_adgroup_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 1111
    Results   []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}
