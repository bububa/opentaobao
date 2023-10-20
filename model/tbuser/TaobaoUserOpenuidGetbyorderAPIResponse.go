package tbuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserOpenuidGetbyorderAPIResponse 根据订单获取买家openuid API返回值
// taobao.user.openuid.getbyorder
//
// 根据订单获取买家openuid，最大查询30个
type TaobaoUserOpenuidGetbyorderAPIResponse struct {
	model.CommonResponse
	TaobaoUserOpenuidGetbyorderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUserOpenuidGetbyorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserOpenuidGetbyorderAPIResponseModel).Reset()
}

// TaobaoUserOpenuidGetbyorderAPIResponseModel is 根据订单获取买家openuid 成功返回结果
type TaobaoUserOpenuidGetbyorderAPIResponseModel struct {
	XMLName xml.Name `xml:"user_openuid_getbyorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 买家uid对象
	OpenUids []OpenUidInfo `json:"open_uids,omitempty" xml:"open_uids>open_uid_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserOpenuidGetbyorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUids = m.OpenUids[:0]
}

var poolTaobaoUserOpenuidGetbyorderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserOpenuidGetbyorderAPIResponse)
	},
}

// GetTaobaoUserOpenuidGetbyorderAPIResponse 从 sync.Pool 获取 TaobaoUserOpenuidGetbyorderAPIResponse
func GetTaobaoUserOpenuidGetbyorderAPIResponse() *TaobaoUserOpenuidGetbyorderAPIResponse {
	return poolTaobaoUserOpenuidGetbyorderAPIResponse.Get().(*TaobaoUserOpenuidGetbyorderAPIResponse)
}

// ReleaseTaobaoUserOpenuidGetbyorderAPIResponse 将 TaobaoUserOpenuidGetbyorderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserOpenuidGetbyorderAPIResponse(v *TaobaoUserOpenuidGetbyorderAPIResponse) {
	v.Reset()
	poolTaobaoUserOpenuidGetbyorderAPIResponse.Put(v)
}
