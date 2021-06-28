package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点容量 APIResponse
tmall.servicecenter.servicestore.deleteservicestorecapacity

删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
必选字段：serviceStoreCode、bizType
*/
type TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse struct {
    model.CommonResponse
    TmallServicecenterServicestoreDeleteservicestorecapacityResponse
}

type TmallServicecenterServicestoreDeleteservicestorecapacityResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_servicestore_deleteservicestorecapacity_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultBase `json:"result,omitempty" xml:"result,omitempty"`

    
}
