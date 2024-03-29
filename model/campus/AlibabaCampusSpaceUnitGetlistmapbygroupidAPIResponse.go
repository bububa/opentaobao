package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse 新增查询多个分组ID各自相关的空间单元信息 API返回值
// alibaba.campus.space.unit.getlistmapbygroupid
//
// 新增查询多个分组ID各自相关的空间单元信息
// HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：	getListMapByGroupIds
type AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponseModel is 新增查询多个分组ID各自相关的空间单元信息 成功返回结果
type AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistmapbygroupid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaCampusSpaceUnitGetlistmapbygroupidMapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse
func GetAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse() *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse {
	return poolAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse.Get().(*AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse 将 AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse(v *AlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistmapbygroupidAPIResponse.Put(v)
}
