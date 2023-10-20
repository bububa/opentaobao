package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopIpoutGetAPIResponse 获取开放平台出口IP段 API返回值
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
type TaobaoTopIpoutGetAPIResponse struct {
	model.CommonResponse
	TaobaoTopIpoutGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopIpoutGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopIpoutGetAPIResponseModel).Reset()
}

// TaobaoTopIpoutGetAPIResponseModel is 获取开放平台出口IP段 成功返回结果
type TaobaoTopIpoutGetAPIResponseModel struct {
	XMLName xml.Name `xml:"top_ipout_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TOP网关出口IP列表
	IpList string `json:"ip_list,omitempty" xml:"ip_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopIpoutGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IpList = ""
}

var poolTaobaoTopIpoutGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopIpoutGetAPIResponse)
	},
}

// GetTaobaoTopIpoutGetAPIResponse 从 sync.Pool 获取 TaobaoTopIpoutGetAPIResponse
func GetTaobaoTopIpoutGetAPIResponse() *TaobaoTopIpoutGetAPIResponse {
	return poolTaobaoTopIpoutGetAPIResponse.Get().(*TaobaoTopIpoutGetAPIResponse)
}

// ReleaseTaobaoTopIpoutGetAPIResponse 将 TaobaoTopIpoutGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopIpoutGetAPIResponse(v *TaobaoTopIpoutGetAPIResponse) {
	v.Reset()
	poolTaobaoTopIpoutGetAPIResponse.Put(v)
}
