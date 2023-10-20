package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberPointOperateAPIResponse 积分操作 API返回值
// alibaba.member.point.operate
//
// 消费会员积分
type AlibabaMemberPointOperateAPIResponse struct {
	model.CommonResponse
	AlibabaMemberPointOperateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberPointOperateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberPointOperateAPIResponseModel).Reset()
}

// AlibabaMemberPointOperateAPIResponseModel is 积分操作 成功返回结果
type AlibabaMemberPointOperateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_point_operate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaMemberPointOperateTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberPointOperateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberPointOperateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberPointOperateAPIResponse)
	},
}

// GetAlibabaMemberPointOperateAPIResponse 从 sync.Pool 获取 AlibabaMemberPointOperateAPIResponse
func GetAlibabaMemberPointOperateAPIResponse() *AlibabaMemberPointOperateAPIResponse {
	return poolAlibabaMemberPointOperateAPIResponse.Get().(*AlibabaMemberPointOperateAPIResponse)
}

// ReleaseAlibabaMemberPointOperateAPIResponse 将 AlibabaMemberPointOperateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberPointOperateAPIResponse(v *AlibabaMemberPointOperateAPIResponse) {
	v.Reset()
	poolAlibabaMemberPointOperateAPIResponse.Put(v)
}
