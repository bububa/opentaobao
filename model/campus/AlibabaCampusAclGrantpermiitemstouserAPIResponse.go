package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemstouserAPIResponse 给人直接授权 API返回值
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
type AlibabaCampusAclGrantpermiitemstouserAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGrantpermiitemstouserAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclGrantpermiitemstouserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclGrantpermiitemstouserAPIResponseModel).Reset()
}

// AlibabaCampusAclGrantpermiitemstouserAPIResponseModel is 给人直接授权 成功返回结果
type AlibabaCampusAclGrantpermiitemstouserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemstouser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclGrantpermiitemstouserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclGrantpermiitemstouserAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclGrantpermiitemstouserAPIResponse)
	},
}

// GetAlibabaCampusAclGrantpermiitemstouserAPIResponse 从 sync.Pool 获取 AlibabaCampusAclGrantpermiitemstouserAPIResponse
func GetAlibabaCampusAclGrantpermiitemstouserAPIResponse() *AlibabaCampusAclGrantpermiitemstouserAPIResponse {
	return poolAlibabaCampusAclGrantpermiitemstouserAPIResponse.Get().(*AlibabaCampusAclGrantpermiitemstouserAPIResponse)
}

// ReleaseAlibabaCampusAclGrantpermiitemstouserAPIResponse 将 AlibabaCampusAclGrantpermiitemstouserAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclGrantpermiitemstouserAPIResponse(v *AlibabaCampusAclGrantpermiitemstouserAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclGrantpermiitemstouserAPIResponse.Put(v)
}
