package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicestoredeleteservicestorecapacityAPIResponse 删除网点容量 API返回值
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
type TmallservicecenterservicestoredeleteservicestorecapacityAPIResponse struct {
	model.CommonResponse
	TmallservicecenterservicestoredeleteservicestorecapacityAPIResponseModel
}

// TmallservicecenterservicestoredeleteservicestorecapacityAPIResponseModel is 删除网点容量 成功返回结果
type TmallservicecenterservicestoredeleteservicestorecapacityAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_deleteservicestorecapacity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
