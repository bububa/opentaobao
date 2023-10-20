package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubuserInfoUpdateAPIResponse 修改指定账户子账号的基本信息 API返回值
// taobao.subuser.info.update
//
// 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息）
type TaobaoSubuserInfoUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSubuserInfoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubuserInfoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubuserInfoUpdateAPIResponseModel).Reset()
}

// TaobaoSubuserInfoUpdateAPIResponseModel is 修改指定账户子账号的基本信息 成功返回结果
type TaobaoSubuserInfoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"subuser_info_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功 true:操作成功; false:操作失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubuserInfoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoSubuserInfoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubuserInfoUpdateAPIResponse)
	},
}

// GetTaobaoSubuserInfoUpdateAPIResponse 从 sync.Pool 获取 TaobaoSubuserInfoUpdateAPIResponse
func GetTaobaoSubuserInfoUpdateAPIResponse() *TaobaoSubuserInfoUpdateAPIResponse {
	return poolTaobaoSubuserInfoUpdateAPIResponse.Get().(*TaobaoSubuserInfoUpdateAPIResponse)
}

// ReleaseTaobaoSubuserInfoUpdateAPIResponse 将 TaobaoSubuserInfoUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubuserInfoUpdateAPIResponse(v *TaobaoSubuserInfoUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSubuserInfoUpdateAPIResponse.Put(v)
}
