package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdateservicestoreAPIResponse 修改网点信息 API返回值
// tmall.servicecenter.servicestore.updateservicestore
//
// 修改网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点存在才能修改。
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
type TmallServicecenterServicestoreUpdateservicestoreAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreUpdateservicestoreAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdateservicestoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreUpdateservicestoreAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreUpdateservicestoreAPIResponseModel is 修改网点信息 成功返回结果
type TmallServicecenterServicestoreUpdateservicestoreAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_updateservicestore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdateservicestoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreUpdateservicestoreAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreUpdateservicestoreAPIResponse)
	},
}

// GetTmallServicecenterServicestoreUpdateservicestoreAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreUpdateservicestoreAPIResponse
func GetTmallServicecenterServicestoreUpdateservicestoreAPIResponse() *TmallServicecenterServicestoreUpdateservicestoreAPIResponse {
	return poolTmallServicecenterServicestoreUpdateservicestoreAPIResponse.Get().(*TmallServicecenterServicestoreUpdateservicestoreAPIResponse)
}

// ReleaseTmallServicecenterServicestoreUpdateservicestoreAPIResponse 将 TmallServicecenterServicestoreUpdateservicestoreAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreUpdateservicestoreAPIResponse(v *TmallServicecenterServicestoreUpdateservicestoreAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreUpdateservicestoreAPIResponse.Put(v)
}
