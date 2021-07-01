package rail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票标准城市查询 API返回值 
alitrip.rail.ir.division.get

国际火车票提供给代理商用于查询标准城市信息，全部城市数据量209530条，含除中国大陆以外的全部海外区域。
代理商通过分页查询的方式，拉取飞猪平台方全部境外标准城市，用于自身城市与飞猪平台城市的映射。
*/
type AlitripRailIrDivisionGetAPIResponse struct {
    model.CommonResponse
    AlitripRailIrDivisionGetAPIResponseModel
}

// 国际火车票标准城市查询 成功返回结果
type AlitripRailIrDivisionGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_rail_ir_division_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象
    Result   *RailResultList `json:"result,omitempty" xml:"result,omitempty"`
}
