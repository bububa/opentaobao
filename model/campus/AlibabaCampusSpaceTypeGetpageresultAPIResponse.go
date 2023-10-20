package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceTypeGetpageresultAPIResponse 分页查询空间类别接口 API返回值
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
type AlibabaCampusSpaceTypeGetpageresultAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceTypeGetpageresultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceTypeGetpageresultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceTypeGetpageresultAPIResponseModel).Reset()
}

// AlibabaCampusSpaceTypeGetpageresultAPIResponseModel is 分页查询空间类别接口 成功返回结果
type AlibabaCampusSpaceTypeGetpageresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_type_getpageresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceTypeGetpageresultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceTypeGetpageresultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceTypeGetpageresultAPIResponse)
	},
}

// GetAlibabaCampusSpaceTypeGetpageresultAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceTypeGetpageresultAPIResponse
func GetAlibabaCampusSpaceTypeGetpageresultAPIResponse() *AlibabaCampusSpaceTypeGetpageresultAPIResponse {
	return poolAlibabaCampusSpaceTypeGetpageresultAPIResponse.Get().(*AlibabaCampusSpaceTypeGetpageresultAPIResponse)
}

// ReleaseAlibabaCampusSpaceTypeGetpageresultAPIResponse 将 AlibabaCampusSpaceTypeGetpageresultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceTypeGetpageresultAPIResponse(v *AlibabaCampusSpaceTypeGetpageresultAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceTypeGetpageresultAPIResponse.Put(v)
}
