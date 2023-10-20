package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetlistAPIResponse 多条件查询空间分组信息 API返回值
// alibaba.campus.space.group.getlist
//
// 多条件查询空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getList
type AlibabaCampusSpaceGroupGetlistAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGroupGetlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceGroupGetlistAPIResponseModel).Reset()
}

// AlibabaCampusSpaceGroupGetlistAPIResponseModel is 多条件查询空间分组信息 成功返回结果
type AlibabaCampusSpaceGroupGetlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_group_getlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGroupGetlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceGroupGetlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceGroupGetlistAPIResponse)
	},
}

// GetAlibabaCampusSpaceGroupGetlistAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetlistAPIResponse
func GetAlibabaCampusSpaceGroupGetlistAPIResponse() *AlibabaCampusSpaceGroupGetlistAPIResponse {
	return poolAlibabaCampusSpaceGroupGetlistAPIResponse.Get().(*AlibabaCampusSpaceGroupGetlistAPIResponse)
}

// ReleaseAlibabaCampusSpaceGroupGetlistAPIResponse 将 AlibabaCampusSpaceGroupGetlistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetlistAPIResponse(v *AlibabaCampusSpaceGroupGetlistAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetlistAPIResponse.Put(v)
}
