package tmc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcGroupDeleteAPIResponse 删除指定的分组或分组下的用户 API返回值
// taobao.tmc.group.delete
//
// 删除指定的分组或分组下的用户，授权消息使用
type TaobaoTmcGroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoTmcGroupDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTmcGroupDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTmcGroupDeleteAPIResponseModel).Reset()
}

// TaobaoTmcGroupDeleteAPIResponseModel is 删除指定的分组或分组下的用户 成功返回结果
type TaobaoTmcGroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_group_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTmcGroupDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTmcGroupDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTmcGroupDeleteAPIResponse)
	},
}

// GetTaobaoTmcGroupDeleteAPIResponse 从 sync.Pool 获取 TaobaoTmcGroupDeleteAPIResponse
func GetTaobaoTmcGroupDeleteAPIResponse() *TaobaoTmcGroupDeleteAPIResponse {
	return poolTaobaoTmcGroupDeleteAPIResponse.Get().(*TaobaoTmcGroupDeleteAPIResponse)
}

// ReleaseTaobaoTmcGroupDeleteAPIResponse 将 TaobaoTmcGroupDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTmcGroupDeleteAPIResponse(v *TaobaoTmcGroupDeleteAPIResponse) {
	v.Reset()
	poolTaobaoTmcGroupDeleteAPIResponse.Put(v)
}
