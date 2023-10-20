package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGetmenubyempidAPIResponse 查询用户的菜单 API返回值
// alibaba.campus.acl.getmenubyempid
//
// 查询用户的菜单
type AlibabaCampusAclGetmenubyempidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGetmenubyempidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclGetmenubyempidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclGetmenubyempidAPIResponseModel).Reset()
}

// AlibabaCampusAclGetmenubyempidAPIResponseModel is 查询用户的菜单 成功返回结果
type AlibabaCampusAclGetmenubyempidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_getmenubyempid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclGetmenubyempidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclGetmenubyempidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclGetmenubyempidAPIResponse)
	},
}

// GetAlibabaCampusAclGetmenubyempidAPIResponse 从 sync.Pool 获取 AlibabaCampusAclGetmenubyempidAPIResponse
func GetAlibabaCampusAclGetmenubyempidAPIResponse() *AlibabaCampusAclGetmenubyempidAPIResponse {
	return poolAlibabaCampusAclGetmenubyempidAPIResponse.Get().(*AlibabaCampusAclGetmenubyempidAPIResponse)
}

// ReleaseAlibabaCampusAclGetmenubyempidAPIResponse 将 AlibabaCampusAclGetmenubyempidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclGetmenubyempidAPIResponse(v *AlibabaCampusAclGetmenubyempidAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclGetmenubyempidAPIResponse.Put(v)
}
