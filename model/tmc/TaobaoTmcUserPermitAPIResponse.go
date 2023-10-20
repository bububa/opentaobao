package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserPermitAPIResponse 为已授权的用户开通消息服务 API返回值
// taobao.tmc.user.permit
//
// 为已授权的用户开通消息服务，授权消息使用。&lt;br/&gt;&lt;span style=&#34;color:red&#34;&gt;注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic&lt;/span&gt;
type TaobaoTmcUserPermitAPIResponse struct {
	model.CommonResponse
	TaobaoTmcUserPermitAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcUserPermitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcUserPermitAPIResponseModel).Reset()
}

// TaobaoTmcUserPermitAPIResponseModel is 为已授权的用户开通消息服务 成功返回结果
type TaobaoTmcUserPermitAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_permit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcUserPermitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTmcUserPermitAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcUserPermitAPIResponse)
	},
}

// GetTaobaoTmcUserPermitAPIResponse 从 sync.Pool 获取 TaobaoTmcUserPermitAPIResponse
func GetTaobaoTmcUserPermitAPIResponse() *TaobaoTmcUserPermitAPIResponse {
	return poolTaobaoTmcUserPermitAPIResponse.Get().(*TaobaoTmcUserPermitAPIResponse)
}

// ReleaseTaobaoTmcUserPermitAPIResponse 将 TaobaoTmcUserPermitAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcUserPermitAPIResponse(v *TaobaoTmcUserPermitAPIResponse) {
	v.Reset()
	poolTaobaoTmcUserPermitAPIResponse.Put(v)
}
