package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse 删除网点容量 API返回值
// tmall.servicecenter.servicestore.deleteservicestorecapacity
//
// 删除网点覆盖的服务，无法恢复，如果只是暂时屏蔽网点的某个能力，可以将此能力对应的网点容量的capacity字段更新为0
// 必选字段：serviceStoreCode、bizType
type TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponseModel is 删除网点容量 成功返回结果
type TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_deleteservicestorecapacity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse)
	},
}

// GetTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse
func GetTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse() *TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse {
	return poolTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse.Get().(*TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse)
}

// ReleaseTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse 将 TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse(v *TmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreDeleteservicestorecapacityAPIResponse.Put(v)
}
