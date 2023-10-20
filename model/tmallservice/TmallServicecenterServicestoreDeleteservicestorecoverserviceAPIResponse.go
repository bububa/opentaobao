package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse 删除网点覆盖的服务 API返回值
// tmall.servicecenter.servicestore.deleteservicestorecoverservice
//
// 天猫服务平台删除网点覆盖的服务，
// 必选字段：serviceStoreCode、bizType
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel is 删除网点覆盖的服务 成功返回结果
type TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_deleteservicestorecoverservice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse)
	},
}

// GetTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse
func GetTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse() *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse {
	return poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse.Get().(*TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse)
}

// ReleaseTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse 将 TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse(v *TmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreDeleteservicestorecoverserviceAPIResponse.Put(v)
}
