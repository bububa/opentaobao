package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGetbyidsAPIResponse 根据ids和类型查询空间列表 API返回值
// alibaba.campus.space.getbyids
//
// 根据ids和类型查询空间列表
type AlibabaCampusSpaceGetbyidsAPIResponse struct {
	model.CommonResponse
	AlibabaCampusSpaceGetbyidsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGetbyidsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusSpaceGetbyidsAPIResponseModel).Reset()
}

// AlibabaCampusSpaceGetbyidsAPIResponseModel is 根据ids和类型查询空间列表 成功返回结果
type AlibabaCampusSpaceGetbyidsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_space_getbyids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外卖订单查询结果
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusSpaceGetbyidsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusSpaceGetbyidsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusSpaceGetbyidsAPIResponse)
	},
}

// GetAlibabaCampusSpaceGetbyidsAPIResponse 从 sync.Pool 获取 AlibabaCampusSpaceGetbyidsAPIResponse
func GetAlibabaCampusSpaceGetbyidsAPIResponse() *AlibabaCampusSpaceGetbyidsAPIResponse {
	return poolAlibabaCampusSpaceGetbyidsAPIResponse.Get().(*AlibabaCampusSpaceGetbyidsAPIResponse)
}

// ReleaseAlibabaCampusSpaceGetbyidsAPIResponse 将 AlibabaCampusSpaceGetbyidsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusSpaceGetbyidsAPIResponse(v *AlibabaCampusSpaceGetbyidsAPIResponse) {
	v.Reset()
	poolAlibabaCampusSpaceGetbyidsAPIResponse.Put(v)
}
