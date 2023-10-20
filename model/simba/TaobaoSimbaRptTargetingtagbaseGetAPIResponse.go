package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbarpttargetingtagbasegetAPIResponse 定向基础报表 API返回值
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
type TaobaosimbarpttargetingtagbasegetAPIResponse struct {
	model.CommonResponse
	TaobaosimbarpttargetingtagbasegetAPIResponseModel
}

// TaobaosimbarpttargetingtagbasegetAPIResponseModel is 定向基础报表 成功返回结果
type TaobaosimbarpttargetingtagbasegetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_targetingtagbase_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []RptBaseEntityDto `json:"results,omitempty" xml:"results>rpt_base_entity_dto,omitempty"`
}
