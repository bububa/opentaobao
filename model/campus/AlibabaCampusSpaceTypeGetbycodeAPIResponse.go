package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceTypeGetbycodeAPIResponse 根据类别编码查询类别 API返回值
// alibaba.campus.space.type.getbycode
//
// 根据类别编码查询类别
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getByCode
type AlibabaCampusSpaceTypeGetbycodeAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceTypeGetbycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceTypeGetbycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceTypeGetbycodeAPIResponseModel).Reset()
}

// AlibabaCampusSpaceTypeGetbycodeAPIResponseModel is 根据类别编码查询类别 成功返回结果
type AlibabaCampusSpaceTypeGetbycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_type_getbycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceTypeGetbycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceTypeGetbycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceTypeGetbycodeAPIResponse)
	},
}

// GetAlibabaCampusSpaceTypeGetbycodeAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceTypeGetbycodeAPIResponse
func GetAlibabaCampusSpaceTypeGetbycodeAPIResponse() *AlibabaCampusSpaceTypeGetbycodeAPIResponse {
	return poolAlibabaCampusSpaceTypeGetbycodeAPIResponse.Get().(*AlibabaCampusSpaceTypeGetbycodeAPIResponse)
}

// ReleaseAlibabaCampusSpaceTypeGetbycodeAPIResponse 将 AlibabaCampusSpaceTypeGetbycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceTypeGetbycodeAPIResponse(v *AlibabaCampusSpaceTypeGetbycodeAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceTypeGetbycodeAPIResponse.Put(v)
}
