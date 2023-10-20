package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse 根据分组ID查询相应的空间单元 API返回值
// alibaba.campus.space.unit.getlistbygroupid
//
// 根据分组ID查询相应的空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByGroupId
type AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel is 根据分组ID查询相应的空间单元 成功返回结果
type AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_unit_getlistbygroupid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceUnitGetlistbygroupidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse)
	},
}

// GetAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse
func GetAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse() *AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse {
	return poolAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse.Get().(*AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse)
}

// ReleaseAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse 将 AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse(v *AlibabaCampusSpaceUnitGetlistbygroupidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistbygroupidAPIResponse.Put(v)
}
