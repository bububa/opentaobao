package crm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGroupMoveAPIResponse 分组移动 API返回值
// taobao.crm.group.move
//
// 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除&lt;br/&gt;注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
type TaobaoCrmGroupMoveAPIResponse struct {
	model.CommonResponse
	TaobaoCrmGroupMoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCrmGroupMoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCrmGroupMoveAPIResponseModel).Reset()
}

// TaobaoCrmGroupMoveAPIResponseModel is 分组移动 成功返回结果
type TaobaoCrmGroupMoveAPIResponseModel struct {
	XMLName xml.Name `xml:"crm_group_move_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCrmGroupMoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoCrmGroupMoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCrmGroupMoveAPIResponse)
	},
}

// GetTaobaoCrmGroupMoveAPIResponse 从 sync.Pool 获取 TaobaoCrmGroupMoveAPIResponse
func GetTaobaoCrmGroupMoveAPIResponse() *TaobaoCrmGroupMoveAPIResponse {
	return poolTaobaoCrmGroupMoveAPIResponse.Get().(*TaobaoCrmGroupMoveAPIResponse)
}

// ReleaseTaobaoCrmGroupMoveAPIResponse 将 TaobaoCrmGroupMoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCrmGroupMoveAPIResponse(v *TaobaoCrmGroupMoveAPIResponse) {
	v.Reset()
	poolTaobaoCrmGroupMoveAPIResponse.Put(v)
}
