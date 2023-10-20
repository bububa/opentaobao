package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse 新增网点覆盖的服务 API返回值
// tmall.servicecenter.servicestore.createservicestorecoverservice
//
// 新增网点覆盖的服务，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点覆盖的服务已存在，会新增失败。
// 网点覆盖的服务包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、授权的类目和品牌
type TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponseModel is 新增网点覆盖的服务 成功返回结果
type TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_createservicestorecoverservice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse)
	},
}

// GetTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse
func GetTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse() *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse {
	return poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse.Get().(*TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse)
}

// ReleaseTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse 将 TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse(v *TmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateservicestorecoverserviceAPIResponse.Put(v)
}
