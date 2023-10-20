package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsHluserUpdateAPIResponse 订单全链路用户信息修改 API返回值
// taobao.jds.hluser.update
//
// 订单全链路用户信息修改，比如是否开放买家端展示
type TaobaoJdsHluserUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoJdsHluserUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsHluserUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsHluserUpdateAPIResponseModel).Reset()
}

// TaobaoJdsHluserUpdateAPIResponseModel is 订单全链路用户信息修改 成功返回结果
type TaobaoJdsHluserUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_hluser_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsHluserUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoJdsHluserUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsHluserUpdateAPIResponse)
	},
}

// GetTaobaoJdsHluserUpdateAPIResponse 从 sync.Pool 获取 TaobaoJdsHluserUpdateAPIResponse
func GetTaobaoJdsHluserUpdateAPIResponse() *TaobaoJdsHluserUpdateAPIResponse {
	return poolTaobaoJdsHluserUpdateAPIResponse.Get().(*TaobaoJdsHluserUpdateAPIResponse)
}

// ReleaseTaobaoJdsHluserUpdateAPIResponse 将 TaobaoJdsHluserUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsHluserUpdateAPIResponse(v *TaobaoJdsHluserUpdateAPIResponse) {
	v.Reset()
	poolTaobaoJdsHluserUpdateAPIResponse.Put(v)
}
