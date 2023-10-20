package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse 空间分组id查业务属性实例 API返回值
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel).Reset()
}

// AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel is 空间分组id查业务属性实例 成功返回结果
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getspacegroupwithattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse)
	},
}

// GetAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse
func GetAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse() *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse {
	return poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse.Get().(*AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse)
}

// ReleaseAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse 将 AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse(v *AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse.Put(v)
}
