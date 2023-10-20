package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinvaccinatecompleteAPIResponse 接种完成反馈接口 API返回值
// alibaba.health.vaccin.vaccinate.complete
//
// ISV 将用户完成接种的疫苗同步给免疫规划中心
type AlibabahealthvaccinvaccinatecompleteAPIResponse struct {
	model.CommonResponse
	AlibabahealthvaccinvaccinatecompleteAPIResponseModel
}

// AlibabahealthvaccinvaccinatecompleteAPIResponseModel is 接种完成反馈接口 成功返回结果
type AlibabahealthvaccinvaccinatecompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_vaccin_vaccinate_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 数据结果实体
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功执行
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
