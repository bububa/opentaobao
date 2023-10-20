package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingAPIResponse 楼栋同步 API返回值
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
type AlibabaAlihouseNewhomeProjectBuildingAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel is 楼栋同步 成功返回结果
type AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_building_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectBuildingResult `json:"result,omitempty" xml:"result,omitempty"`
}
