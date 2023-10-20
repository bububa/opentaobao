package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceAttrSetattrAPIResponse 新增业务属性实例接口 API返回值
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
type AlibabaCampusSpaceAttrSetattrAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceAttrSetattrAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceAttrSetattrAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceAttrSetattrAPIResponseModel).Reset()
}

// AlibabaCampusSpaceAttrSetattrAPIResponseModel is 新增业务属性实例接口 成功返回结果
type AlibabaCampusSpaceAttrSetattrAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_attr_setattr_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceAttrSetattrAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceAttrSetattrAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceAttrSetattrAPIResponse)
	},
}

// GetAlibabaCampusSpaceAttrSetattrAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceAttrSetattrAPIResponse
func GetAlibabaCampusSpaceAttrSetattrAPIResponse() *AlibabaCampusSpaceAttrSetattrAPIResponse {
	return poolAlibabaCampusSpaceAttrSetattrAPIResponse.Get().(*AlibabaCampusSpaceAttrSetattrAPIResponse)
}

// ReleaseAlibabaCampusSpaceAttrSetattrAPIResponse 将 AlibabaCampusSpaceAttrSetattrAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceAttrSetattrAPIResponse(v *AlibabaCampusSpaceAttrSetattrAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceAttrSetattrAPIResponse.Put(v)
}
