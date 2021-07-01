package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse
删除网点覆盖的服务 API返回值
tmall.servicecenter.servicestore.deleteservicestorecoverservice

天猫服务平台删除网点覆盖的服务，
必选字段：serviceStoreCode、bizType */
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel
}

// TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel is 删除网点覆盖的服务 成功返回结果
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_deleteservicestorecoverservice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}
