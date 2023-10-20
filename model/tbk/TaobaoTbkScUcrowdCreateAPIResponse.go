package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdCreateAPIResponse 淘宝客-服务商-创建人群标签 API返回值
// taobao.tbk.sc.ucrowd.create
//
// 服务商使用。可为淘宝客创建会员人群标签，获得人群id，人群id可用于物料评估推荐等场景。
type TaobaoTbkScUcrowdCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScUcrowdCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScUcrowdCreateAPIResponseModel).Reset()
}

// TaobaoTbkScUcrowdCreateAPIResponseModel is 淘宝客-服务商-创建人群标签 成功返回结果
type TaobaoTbkScUcrowdCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdCreateRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScUcrowdCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdCreateAPIResponse)
	},
}

// GetTaobaoTbkScUcrowdCreateAPIResponse 从 sync.Pool 获取 TaobaoTbkScUcrowdCreateAPIResponse
func GetTaobaoTbkScUcrowdCreateAPIResponse() *TaobaoTbkScUcrowdCreateAPIResponse {
	return poolTaobaoTbkScUcrowdCreateAPIResponse.Get().(*TaobaoTbkScUcrowdCreateAPIResponse)
}

// ReleaseTaobaoTbkScUcrowdCreateAPIResponse 将 TaobaoTbkScUcrowdCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScUcrowdCreateAPIResponse(v *TaobaoTbkScUcrowdCreateAPIResponse) {
	v.Reset()
	poolTaobaoTbkScUcrowdCreateAPIResponse.Put(v)
}
