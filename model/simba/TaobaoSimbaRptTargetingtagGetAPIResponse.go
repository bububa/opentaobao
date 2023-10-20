package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarpttargetingtaggetAPIResponse 搜索人群离线报表 API返回值
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
type TaobaosimbarpttargetingtaggetAPIResponse struct {
	model.CommonResponse
	TaobaosimbarpttargetingtaggetAPIResponseModel
}

// TaobaosimbarpttargetingtaggetAPIResponseModel is 搜索人群离线报表 成功返回结果
type TaobaosimbarpttargetingtaggetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_targetingtag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 111
	Results []RtRptResultEntityDto `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}
