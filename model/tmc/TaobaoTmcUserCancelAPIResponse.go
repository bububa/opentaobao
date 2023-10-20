package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserCancelAPIResponse 取消用户的消息服务 API返回值
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
type TaobaoTmcUserCancelAPIResponse struct {
	model.CommonResponse
	TaobaoTmcUserCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcUserCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcUserCancelAPIResponseModel).Reset()
}

// TaobaoTmcUserCancelAPIResponseModel is 取消用户的消息服务 成功返回结果
type TaobaoTmcUserCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功,如果为false并且没有错误码，表示删除的用户不存在。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcUserCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTmcUserCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcUserCancelAPIResponse)
	},
}

// GetTaobaoTmcUserCancelAPIResponse 从 sync.Pool 获取 TaobaoTmcUserCancelAPIResponse
func GetTaobaoTmcUserCancelAPIResponse() *TaobaoTmcUserCancelAPIResponse {
	return poolTaobaoTmcUserCancelAPIResponse.Get().(*TaobaoTmcUserCancelAPIResponse)
}

// ReleaseTaobaoTmcUserCancelAPIResponse 将 TaobaoTmcUserCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcUserCancelAPIResponse(v *TaobaoTmcUserCancelAPIResponse) {
	v.Reset()
	poolTaobaoTmcUserCancelAPIResponse.Put(v)
}
