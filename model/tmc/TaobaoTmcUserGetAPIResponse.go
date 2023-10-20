package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserGetAPIResponse 获取用户已开通消息 API返回值
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetAPIResponse struct {
	model.CommonResponse
	TaobaoTmcUserGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcUserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcUserGetAPIResponseModel).Reset()
}

// TaobaoTmcUserGetAPIResponseModel is 获取用户已开通消息 成功返回结果
type TaobaoTmcUserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开通的用户数据
	TmcUser *TmcUser `json:"tmc_user,omitempty" xml:"tmc_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcUserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmcUser = nil
}

var poolTaobaoTmcUserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcUserGetAPIResponse)
	},
}

// GetTaobaoTmcUserGetAPIResponse 从 sync.Pool 获取 TaobaoTmcUserGetAPIResponse
func GetTaobaoTmcUserGetAPIResponse() *TaobaoTmcUserGetAPIResponse {
	return poolTaobaoTmcUserGetAPIResponse.Get().(*TaobaoTmcUserGetAPIResponse)
}

// ReleaseTaobaoTmcUserGetAPIResponse 将 TaobaoTmcUserGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcUserGetAPIResponse(v *TaobaoTmcUserGetAPIResponse) {
	v.Reset()
	poolTaobaoTmcUserGetAPIResponse.Put(v)
}
