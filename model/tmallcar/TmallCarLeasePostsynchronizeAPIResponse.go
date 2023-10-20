package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeasePostsynchronizeAPIResponse 天猫开新车租后信息同步 API返回值
// tmall.car.lease.postsynchronize
//
// 商家同步天猫开新车租后方案
type TmallCarLeasePostsynchronizeAPIResponse struct {
	model.CommonResponse
	TmallCarLeasePostsynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeasePostsynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeasePostsynchronizeAPIResponseModel).Reset()
}

// TmallCarLeasePostsynchronizeAPIResponseModel is 天猫开新车租后信息同步 成功返回结果
type TmallCarLeasePostsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_postsynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeasePostsynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeasePostsynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeasePostsynchronizeAPIResponse)
	},
}

// GetTmallCarLeasePostsynchronizeAPIResponse 从 sync.Pool 获取 TmallCarLeasePostsynchronizeAPIResponse
func GetTmallCarLeasePostsynchronizeAPIResponse() *TmallCarLeasePostsynchronizeAPIResponse {
	return poolTmallCarLeasePostsynchronizeAPIResponse.Get().(*TmallCarLeasePostsynchronizeAPIResponse)
}

// ReleaseTmallCarLeasePostsynchronizeAPIResponse 将 TmallCarLeasePostsynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeasePostsynchronizeAPIResponse(v *TmallCarLeasePostsynchronizeAPIResponse) {
	v.Reset()
	poolTmallCarLeasePostsynchronizeAPIResponse.Put(v)
}
