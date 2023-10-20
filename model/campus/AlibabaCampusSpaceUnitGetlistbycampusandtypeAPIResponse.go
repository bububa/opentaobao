package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse 根据园区id及TypeId获取空间单元 API返回值
// alibaba.campus.space.unit.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByCampusAndType
type AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponseModel is 根据园区id及TypeId获取空间单元 成功返回结果
type AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistbycampusandtype_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse
func GetAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse() *AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse {
	return poolAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse.Get().(*AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse 将 AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse(v *AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse.Put(v)
}
