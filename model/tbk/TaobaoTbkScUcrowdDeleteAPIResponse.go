package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdDeleteAPIResponse 淘宝客-服务商-删除人群标签 API返回值
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
type TaobaoTbkScUcrowdDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScUcrowdDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScUcrowdDeleteAPIResponseModel).Reset()
}

// TaobaoTbkScUcrowdDeleteAPIResponseModel is 淘宝客-服务商-删除人群标签 成功返回结果
type TaobaoTbkScUcrowdDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdDeleteRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScUcrowdDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdDeleteAPIResponse)
	},
}

// GetTaobaoTbkScUcrowdDeleteAPIResponse 从 sync.Pool 获取 TaobaoTbkScUcrowdDeleteAPIResponse
func GetTaobaoTbkScUcrowdDeleteAPIResponse() *TaobaoTbkScUcrowdDeleteAPIResponse {
	return poolTaobaoTbkScUcrowdDeleteAPIResponse.Get().(*TaobaoTbkScUcrowdDeleteAPIResponse)
}

// ReleaseTaobaoTbkScUcrowdDeleteAPIResponse 将 TaobaoTbkScUcrowdDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScUcrowdDeleteAPIResponse(v *TaobaoTbkScUcrowdDeleteAPIResponse) {
	v.Reset()
	poolTaobaoTbkScUcrowdDeleteAPIResponse.Put(v)
}
