package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse 新增网点容量 API返回值
// tmall.servicecenter.servicestore.createservicestorecapacity
//
// 新增网点容量，唯一性校验：服务商淘宝账号+网点编码+biz_type
// 前提是网点要存在，
// 如果需要新增的网点容量已存在，会新增失败。
// 网点容量包含了业务类型(比如电器预约安装)、天猫服务的servicecode列表、类目区域和容量
type TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreCreateservicestorecapacityAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreCreateservicestorecapacityAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreCreateservicestorecapacityAPIResponseModel is 新增网点容量 成功返回结果
type TmallServicecenterServicestoreCreateservicestorecapacityAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_createservicestorecapacity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateservicestorecapacityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse)
	},
}

// GetTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse
func GetTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse() *TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse {
	return poolTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse.Get().(*TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse)
}

// ReleaseTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse 将 TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse(v *TmallServicecenterServicestoreCreateservicestorecapacityAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateservicestorecapacityAPIResponse.Put(v)
}
