package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceCampusGetbyidAPIResponse 根据园区id获取园区信息 API返回值
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
type AlibabaCampusSpaceCampusGetbyidAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceCampusGetbyidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceCampusGetbyidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceCampusGetbyidAPIResponseModel).Reset()
}

// AlibabaCampusSpaceCampusGetbyidAPIResponseModel is 根据园区id获取园区信息 成功返回结果
type AlibabaCampusSpaceCampusGetbyidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_campus_getbyid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceCampusGetbyidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceCampusGetbyidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceCampusGetbyidAPIResponse)
	},
}

// GetAlibabaCampusSpaceCampusGetbyidAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceCampusGetbyidAPIResponse
func GetAlibabaCampusSpaceCampusGetbyidAPIResponse() *AlibabaCampusSpaceCampusGetbyidAPIResponse {
	return poolAlibabaCampusSpaceCampusGetbyidAPIResponse.Get().(*AlibabaCampusSpaceCampusGetbyidAPIResponse)
}

// ReleaseAlibabaCampusSpaceCampusGetbyidAPIResponse 将 AlibabaCampusSpaceCampusGetbyidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceCampusGetbyidAPIResponse(v *AlibabaCampusSpaceCampusGetbyidAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceCampusGetbyidAPIResponse.Put(v)
}
