package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse 根据园区id及TypeId获取空间分组 API返回值
// alibaba.campus.space.group.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间分组
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getListByCampusAndType
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponseModel).Reset()
}

// AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponseModel is 根据园区id及TypeId获取空间分组 成功返回结果
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getlistbycampusandtype_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse)
	},
}

// GetAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse
func GetAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse() *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse {
	return poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse.Get().(*AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse)
}

// ReleaseAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse 将 AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse(v *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse.Put(v)
}
