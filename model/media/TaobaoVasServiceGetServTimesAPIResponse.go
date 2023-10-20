package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasServiceGetServTimesAPIResponse 查询某个用户图片空间的使用情况 API返回值
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
type TaobaoVasServiceGetServTimesAPIResponse struct {
	model.CommonResponse
	TaobaoVasServiceGetServTimesAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVasServiceGetServTimesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVasServiceGetServTimesAPIResponseModel).Reset()
}

// TaobaoVasServiceGetServTimesAPIResponseModel is 查询某个用户图片空间的使用情况 成功返回结果
type TaobaoVasServiceGetServTimesAPIResponseModel struct {
	XMLName xml.Name `xml:"vas_service_getServTimes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 总次数（容量）
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 剩余次数（容量）
	LeftNum int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVasServiceGetServTimesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TotalNum = 0
	m.LeftNum = 0
}

var poolTaobaoVasServiceGetServTimesAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVasServiceGetServTimesAPIResponse)
	},
}

// GetTaobaoVasServiceGetServTimesAPIResponse 从 sync.Pool 获取 TaobaoVasServiceGetServTimesAPIResponse
func GetTaobaoVasServiceGetServTimesAPIResponse() *TaobaoVasServiceGetServTimesAPIResponse {
	return poolTaobaoVasServiceGetServTimesAPIResponse.Get().(*TaobaoVasServiceGetServTimesAPIResponse)
}

// ReleaseTaobaoVasServiceGetServTimesAPIResponse 将 TaobaoVasServiceGetServTimesAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVasServiceGetServTimesAPIResponse(v *TaobaoVasServiceGetServTimesAPIResponse) {
	v.Reset()
	poolTaobaoVasServiceGetServTimesAPIResponse.Put(v)
}
