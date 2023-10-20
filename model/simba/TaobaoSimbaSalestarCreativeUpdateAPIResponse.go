package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarCreativeUpdateAPIResponse 销量明星更新创意相关接口 API返回值
// taobao.simba.salestar.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
type TaobaoSimbaSalestarCreativeUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarCreativeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarCreativeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarCreativeUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarCreativeUpdateAPIResponseModel is 销量明星更新创意相关接口 成功返回结果
type TaobaoSimbaSalestarCreativeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_creative_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创意修改记录对象
	Creativerecord *CreativeRecord `json:"creativerecord,omitempty" xml:"creativerecord,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarCreativeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Creativerecord = nil
}

var poolTaobaoSimbaSalestarCreativeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarCreativeUpdateAPIResponse)
	},
}

// GetTaobaoSimbaSalestarCreativeUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarCreativeUpdateAPIResponse
func GetTaobaoSimbaSalestarCreativeUpdateAPIResponse() *TaobaoSimbaSalestarCreativeUpdateAPIResponse {
	return poolTaobaoSimbaSalestarCreativeUpdateAPIResponse.Get().(*TaobaoSimbaSalestarCreativeUpdateAPIResponse)
}

// ReleaseTaobaoSimbaSalestarCreativeUpdateAPIResponse 将 TaobaoSimbaSalestarCreativeUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarCreativeUpdateAPIResponse(v *TaobaoSimbaSalestarCreativeUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarCreativeUpdateAPIResponse.Put(v)
}
