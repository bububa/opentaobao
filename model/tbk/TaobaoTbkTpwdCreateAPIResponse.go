package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkTpwdCreateAPIResponse 淘宝客-公用-淘口令生成 API返回值
// taobao.tbk.tpwd.create
//
// 提供淘口令生成接口。提交需要生成淘口令的内容、logo、url等参数，生成淘口令，其中关键信息为￥SADadW￥，后续可基于淘口令进行文案包装组装用于传播。
type TaobaoTbkTpwdCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTbkTpwdCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkTpwdCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkTpwdCreateAPIResponseModel).Reset()
}

// TaobaoTbkTpwdCreateAPIResponseModel is 淘宝客-公用-淘口令生成 成功返回结果
type TaobaoTbkTpwdCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_tpwd_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Data *TaobaoTbkTpwdCreateMapData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkTpwdCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkTpwdCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkTpwdCreateAPIResponse)
	},
}

// GetTaobaoTbkTpwdCreateAPIResponse 从 sync.Pool 获取 TaobaoTbkTpwdCreateAPIResponse
func GetTaobaoTbkTpwdCreateAPIResponse() *TaobaoTbkTpwdCreateAPIResponse {
	return poolTaobaoTbkTpwdCreateAPIResponse.Get().(*TaobaoTbkTpwdCreateAPIResponse)
}

// ReleaseTaobaoTbkTpwdCreateAPIResponse 将 TaobaoTbkTpwdCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkTpwdCreateAPIResponse(v *TaobaoTbkTpwdCreateAPIResponse) {
	v.Reset()
	poolTaobaoTbkTpwdCreateAPIResponse.Put(v)
}
