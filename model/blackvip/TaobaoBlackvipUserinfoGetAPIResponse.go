package blackvip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBlackvipUserinfoGetAPIResponse 88VIP用户信息查询 API返回值
// taobao.blackvip.userinfo.get
//
// 查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等
type TaobaoBlackvipUserinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoBlackvipUserinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBlackvipUserinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBlackvipUserinfoGetAPIResponseModel).Reset()
}

// TaobaoBlackvipUserinfoGetAPIResponseModel is 88VIP用户信息查询 成功返回结果
type TaobaoBlackvipUserinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"blackvip_userinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果支持对象
	Result *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBlackvipUserinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBlackvipUserinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBlackvipUserinfoGetAPIResponse)
	},
}

// GetTaobaoBlackvipUserinfoGetAPIResponse 从 sync.Pool 获取 TaobaoBlackvipUserinfoGetAPIResponse
func GetTaobaoBlackvipUserinfoGetAPIResponse() *TaobaoBlackvipUserinfoGetAPIResponse {
	return poolTaobaoBlackvipUserinfoGetAPIResponse.Get().(*TaobaoBlackvipUserinfoGetAPIResponse)
}

// ReleaseTaobaoBlackvipUserinfoGetAPIResponse 将 TaobaoBlackvipUserinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBlackvipUserinfoGetAPIResponse(v *TaobaoBlackvipUserinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoBlackvipUserinfoGetAPIResponse.Put(v)
}
