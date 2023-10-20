package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseCitysynchronizeAPIResponse 天猫开新车租后分期城市信息同步 API返回值
// tmall.car.lease.citysynchronize
//
// 天猫开新车租后分期城市信息同步
type TmallCarLeaseCitysynchronizeAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseCitysynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseCitysynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseCitysynchronizeAPIResponseModel).Reset()
}

// TmallCarLeaseCitysynchronizeAPIResponseModel is 天猫开新车租后分期城市信息同步 成功返回结果
type TmallCarLeaseCitysynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_citysynchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseCitysynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseCitysynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseCitysynchronizeAPIResponse)
	},
}

// GetTmallCarLeaseCitysynchronizeAPIResponse 从 sync.Pool 获取 TmallCarLeaseCitysynchronizeAPIResponse
func GetTmallCarLeaseCitysynchronizeAPIResponse() *TmallCarLeaseCitysynchronizeAPIResponse {
	return poolTmallCarLeaseCitysynchronizeAPIResponse.Get().(*TmallCarLeaseCitysynchronizeAPIResponse)
}

// ReleaseTmallCarLeaseCitysynchronizeAPIResponse 将 TmallCarLeaseCitysynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseCitysynchronizeAPIResponse(v *TmallCarLeaseCitysynchronizeAPIResponse) {
	v.Reset()
	poolTmallCarLeaseCitysynchronizeAPIResponse.Put(v)
}
