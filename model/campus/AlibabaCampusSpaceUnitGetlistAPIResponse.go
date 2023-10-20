package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistAPIResponse 多条件查询空间单元信息 API返回值
// alibaba.campus.space.unit.getlist
//
// 多条件查询空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getList
type AlibabaCampusSpaceUnitGetlistAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetlistAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetlistAPIResponseModel is 多条件查询空间单元信息 成功返回结果
type AlibabaCampusSpaceUnitGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// results
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetlistAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetlistAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistAPIResponse
func GetAlibabaCampusSpaceUnitGetlistAPIResponse() *AlibabaCampusSpaceUnitGetlistAPIResponse {
	return poolAlibabaCampusSpaceUnitGetlistAPIResponse.Get().(*AlibabaCampusSpaceUnitGetlistAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetlistAPIResponse 将 AlibabaCampusSpaceUnitGetlistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistAPIResponse(v *AlibabaCampusSpaceUnitGetlistAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistAPIResponse.Put(v)
}
