package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterServiceproductQueryAPIResponse 天猫服务商服务产品查询 API返回值
// tmall.mallitemcenter.serviceproduct.query
//
// 查询天猫服务的服务商发布的服务产品
type TmallMallitemcenterServiceproductQueryAPIResponse struct {
	model.CommonResponse
	TmallMallitemcenterServiceproductQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMallitemcenterServiceproductQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMallitemcenterServiceproductQueryAPIResponseModel).Reset()
}

// TmallMallitemcenterServiceproductQueryAPIResponseModel is 天猫服务商服务产品查询 成功返回结果
type TmallMallitemcenterServiceproductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mallitemcenter_serviceproduct_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallMallitemcenterServiceproductQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallMallitemcenterServiceproductQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallMallitemcenterServiceproductQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMallitemcenterServiceproductQueryAPIResponse)
	},
}

// GetTmallMallitemcenterServiceproductQueryAPIResponse 从 sync.Pool 获取 TmallMallitemcenterServiceproductQueryAPIResponse
func GetTmallMallitemcenterServiceproductQueryAPIResponse() *TmallMallitemcenterServiceproductQueryAPIResponse {
	return poolTmallMallitemcenterServiceproductQueryAPIResponse.Get().(*TmallMallitemcenterServiceproductQueryAPIResponse)
}

// ReleaseTmallMallitemcenterServiceproductQueryAPIResponse 将 TmallMallitemcenterServiceproductQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallMallitemcenterServiceproductQueryAPIResponse(v *TmallMallitemcenterServiceproductQueryAPIResponse) {
	v.Reset()
	poolTmallMallitemcenterServiceproductQueryAPIResponse.Put(v)
}
