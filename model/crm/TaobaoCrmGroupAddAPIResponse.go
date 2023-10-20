package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupAddAPIResponse 卖家创建一个分组 API返回值
// taobao.crm.group.add
//
// 卖家创建一个新的分组，接口返回一个创建成功的分组的id
type TaobaoCrmGroupAddAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGroupAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGroupAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGroupAddAPIResponseModel).Reset()
}

// TaobaoCrmGroupAddAPIResponseModel is 卖家创建一个分组 成功返回结果
type TaobaoCrmGroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增分组的id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 添加分组是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGroupAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GroupId = 0
	m.IsSuccess = false
}

var poolTaobaoCrmGroupAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGroupAddAPIResponse)
	},
}

// GetTaobaoCrmGroupAddAPIResponse 从 sync.Pool 获取 TaobaoCrmGroupAddAPIResponse
func GetTaobaoCrmGroupAddAPIResponse() *TaobaoCrmGroupAddAPIResponse {
	return poolTaobaoCrmGroupAddAPIResponse.Get().(*TaobaoCrmGroupAddAPIResponse)
}

// ReleaseTaobaoCrmGroupAddAPIResponse 将 TaobaoCrmGroupAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGroupAddAPIResponse(v *TaobaoCrmGroupAddAPIResponse) {
	v.Reset()
	poolTaobaoCrmGroupAddAPIResponse.Put(v)
}
