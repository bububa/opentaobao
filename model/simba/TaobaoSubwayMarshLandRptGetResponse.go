package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取捡漏词包分时报表数据 APIResponse
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据
*/
type TaobaoSubwayMarshLandRptGetAPIResponse struct {
    model.CommonResponse
    TaobaoSubwayMarshLandRptGetResponse
}

type TaobaoSubwayMarshLandRptGetResponse struct {
    XMLName xml.Name `xml:"subway_marsh_land_rpt_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 捡漏词包分时报表数据列表
    
    ResultList   []RptResult `json:"result_list,omitempty" xml:"result_list>rpt_result,omitempty"`
    
    
}
