package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponse 新房楼栋修改e码 API返回值
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
type AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponseModel
}

// AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponseModel is 新房楼栋修改e码 成功返回结果
type AlibabaalihousenewhomeprojectbuildingecodeupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_building_ecode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectbuildingecodeupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
