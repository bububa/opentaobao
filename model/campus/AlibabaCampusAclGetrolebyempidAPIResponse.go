package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetrolebyempidAPIResponse 根据用户查询角色 API返回值
// alibaba.campus.acl.getrolebyempid
//
// 根据用户查询角色
type AlibabaCampusAclGetrolebyempidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGetrolebyempidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclGetrolebyempidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclGetrolebyempidAPIResponseModel).Reset()
}

// AlibabaCampusAclGetrolebyempidAPIResponseModel is 根据用户查询角色 成功返回结果
type AlibabaCampusAclGetrolebyempidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_getrolebyempid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclGetrolebyempidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclGetrolebyempidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclGetrolebyempidAPIResponse)
	},
}

// GetAlibabaCampusAclGetrolebyempidAPIResponse 从 sync.Pool 获取 AlibabaCampusAclGetrolebyempidAPIResponse
func GetAlibabaCampusAclGetrolebyempidAPIResponse() *AlibabaCampusAclGetrolebyempidAPIResponse {
	return poolAlibabaCampusAclGetrolebyempidAPIResponse.Get().(*AlibabaCampusAclGetrolebyempidAPIResponse)
}

// ReleaseAlibabaCampusAclGetrolebyempidAPIResponse 将 AlibabaCampusAclGetrolebyempidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclGetrolebyempidAPIResponse(v *AlibabaCampusAclGetrolebyempidAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclGetrolebyempidAPIResponse.Put(v)
}
