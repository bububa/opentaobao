package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsHluserGetAPIResponse 订单全链路用户信息获取 API返回值
// taobao.jds.hluser.get
//
// 订单全链路用户信息获取
type TaobaoJdsHluserGetAPIResponse struct {
	model.CommonResponse
	TaobaoJdsHluserGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsHluserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsHluserGetAPIResponseModel).Reset()
}

// TaobaoJdsHluserGetAPIResponseModel is 订单全链路用户信息获取 成功返回结果
type TaobaoJdsHluserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_hluser_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回流用户信息
	HlUser *HlUserDo `json:"hl_user,omitempty" xml:"hl_user,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsHluserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HlUser = nil
}

var poolTaobaoJdsHluserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsHluserGetAPIResponse)
	},
}

// GetTaobaoJdsHluserGetAPIResponse 从 sync.Pool 获取 TaobaoJdsHluserGetAPIResponse
func GetTaobaoJdsHluserGetAPIResponse() *TaobaoJdsHluserGetAPIResponse {
	return poolTaobaoJdsHluserGetAPIResponse.Get().(*TaobaoJdsHluserGetAPIResponse)
}

// ReleaseTaobaoJdsHluserGetAPIResponse 将 TaobaoJdsHluserGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsHluserGetAPIResponse(v *TaobaoJdsHluserGetAPIResponse) {
	v.Reset()
	poolTaobaoJdsHluserGetAPIResponse.Put(v)
}
