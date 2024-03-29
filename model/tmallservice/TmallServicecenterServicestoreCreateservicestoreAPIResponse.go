package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestoreAPIResponse 服务网点创建 API返回值
// tmall.servicecenter.servicestore.createservicestore
//
// 创建网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点不存在才能创建。地址信息：中文和编码二选一，都填则以编码address_code为准。
// 错误码
// 1, 服务商昵称无效
// 2, 缺少省份
// 3, 缺少城市
// 4, 缺少区域
// 5, 缺少详细地址
// 6, 传入地址不在标准地址库中，无法解析
// 7, 缺少网点编码
// 8, 缺少网点名称
// 9, 缺少网点电话
// 10, 网点已存在
// 11, 网点不存在
// 12, 系统错误
type TmallServicecenterServicestoreCreateservicestoreAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreCreateservicestoreAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreCreateservicestoreAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreCreateservicestoreAPIResponseModel is 服务网点创建 成功返回结果
type TmallServicecenterServicestoreCreateservicestoreAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_createservicestore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreCreateservicestoreAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreCreateservicestoreAPIResponse)
	},
}

// GetTmallServicecenterServicestoreCreateservicestoreAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreCreateservicestoreAPIResponse
func GetTmallServicecenterServicestoreCreateservicestoreAPIResponse() *TmallServicecenterServicestoreCreateservicestoreAPIResponse {
	return poolTmallServicecenterServicestoreCreateservicestoreAPIResponse.Get().(*TmallServicecenterServicestoreCreateservicestoreAPIResponse)
}

// ReleaseTmallServicecenterServicestoreCreateservicestoreAPIResponse 将 TmallServicecenterServicestoreCreateservicestoreAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateservicestoreAPIResponse(v *TmallServicecenterServicestoreCreateservicestoreAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateservicestoreAPIResponse.Put(v)
}
