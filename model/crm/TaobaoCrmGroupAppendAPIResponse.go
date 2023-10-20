package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupAppendAPIResponse 将一个分组添加到另外一个分组 API返回值
// taobao.crm.group.append
//
// 将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。
type TaobaoCrmGroupAppendAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGroupAppendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGroupAppendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGroupAppendAPIResponseModel).Reset()
}

// TaobaoCrmGroupAppendAPIResponseModel is 将一个分组添加到另外一个分组 成功返回结果
type TaobaoCrmGroupAppendAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_append_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步任务请求成功，添加任务是否完成通过taobao.crm.grouptask.check检测
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGroupAppendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmGroupAppendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGroupAppendAPIResponse)
	},
}

// GetTaobaoCrmGroupAppendAPIResponse 从 sync.Pool 获取 TaobaoCrmGroupAppendAPIResponse
func GetTaobaoCrmGroupAppendAPIResponse() *TaobaoCrmGroupAppendAPIResponse {
	return poolTaobaoCrmGroupAppendAPIResponse.Get().(*TaobaoCrmGroupAppendAPIResponse)
}

// ReleaseTaobaoCrmGroupAppendAPIResponse 将 TaobaoCrmGroupAppendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGroupAppendAPIResponse(v *TaobaoCrmGroupAppendAPIResponse) {
	v.Reset()
	poolTaobaoCrmGroupAppendAPIResponse.Put(v)
}
