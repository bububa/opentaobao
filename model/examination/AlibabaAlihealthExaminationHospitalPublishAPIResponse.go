package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationHospitalPublishAPIResponse 体检机构对接_门店发布／更新 API返回值
// alibaba.alihealth.examination.hospital.publish
//
// 第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口
type AlibabaAlihealthExaminationHospitalPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationHospitalPublishAPIResponseModel
}

// AlibabaAlihealthExaminationHospitalPublishAPIResponseModel is 体检机构对接_门店发布／更新 成功返回结果
type AlibabaAlihealthExaminationHospitalPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_hospital_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
