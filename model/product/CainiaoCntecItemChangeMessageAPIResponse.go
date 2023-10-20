package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecItemChangeMessageAPIResponse 商品变更消息 API返回值
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
type CainiaoCntecItemChangeMessageAPIResponse struct {
	model.CommonResponse
	CainiaoCntecItemChangeMessageAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCntecItemChangeMessageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCntecItemChangeMessageAPIResponseModel).Reset()
}

// CainiaoCntecItemChangeMessageAPIResponseModel is 商品变更消息 成功返回结果
type CainiaoCntecItemChangeMessageAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_item_change_message_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回的result结构体
	Result *CainiaoCntecItemChangeMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCntecItemChangeMessageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCntecItemChangeMessageAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCntecItemChangeMessageAPIResponse)
	},
}

// GetCainiaoCntecItemChangeMessageAPIResponse 从 sync.Pool 获取 CainiaoCntecItemChangeMessageAPIResponse
func GetCainiaoCntecItemChangeMessageAPIResponse() *CainiaoCntecItemChangeMessageAPIResponse {
	return poolCainiaoCntecItemChangeMessageAPIResponse.Get().(*CainiaoCntecItemChangeMessageAPIResponse)
}

// ReleaseCainiaoCntecItemChangeMessageAPIResponse 将 CainiaoCntecItemChangeMessageAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCntecItemChangeMessageAPIResponse(v *CainiaoCntecItemChangeMessageAPIResponse) {
	v.Reset()
	poolCainiaoCntecItemChangeMessageAPIResponse.Put(v)
}
