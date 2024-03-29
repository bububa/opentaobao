package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupUpdateAPIResponse 修改一个已经存在的分组 API返回值
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
type TaobaoCrmGroupUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGroupUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGroupUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGroupUpdateAPIResponseModel).Reset()
}

// TaobaoCrmGroupUpdateAPIResponseModel is 修改一个已经存在的分组 成功返回结果
type TaobaoCrmGroupUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组修改是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGroupUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmGroupUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGroupUpdateAPIResponse)
	},
}

// GetTaobaoCrmGroupUpdateAPIResponse 从 sync.Pool 获取 TaobaoCrmGroupUpdateAPIResponse
func GetTaobaoCrmGroupUpdateAPIResponse() *TaobaoCrmGroupUpdateAPIResponse {
	return poolTaobaoCrmGroupUpdateAPIResponse.Get().(*TaobaoCrmGroupUpdateAPIResponse)
}

// ReleaseTaobaoCrmGroupUpdateAPIResponse 将 TaobaoCrmGroupUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGroupUpdateAPIResponse(v *TaobaoCrmGroupUpdateAPIResponse) {
	v.Reset()
	poolTaobaoCrmGroupUpdateAPIResponse.Put(v)
}
