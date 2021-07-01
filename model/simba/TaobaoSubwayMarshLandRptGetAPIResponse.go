package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubwayMarshLandRptGetAPIResponse
获取捡漏词包分时报表数据 API返回值
taobao.subway.marsh.land.rpt.get

获取捡漏词包分时报表数据 */
type TaobaoSubwayMarshLandRptGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayMarshLandRptGetAPIResponseModel
}

// TaobaoSubwayMarshLandRptGetAPIResponseModel is 获取捡漏词包分时报表数据 成功返回结果
type TaobaoSubwayMarshLandRptGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_marsh_land_rpt_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 捡漏词包分时报表数据列表
	ResultList []RptResult `json:"result_list,omitempty" xml:"result_list>rpt_result,omitempty"`
}
