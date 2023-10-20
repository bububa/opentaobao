package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse 分页查询空间分组业务属性 API返回值
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponseModel).Reset()
}

// AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponseModel is 分页查询空间分组业务属性 成功返回结果
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegrouplistwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse)
	},
}

// GetAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse
func GetAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse() *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse {
	return poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse.Get().(*AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse)
}

// ReleaseAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse 将 AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse(v *AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse.Put(v)
}
