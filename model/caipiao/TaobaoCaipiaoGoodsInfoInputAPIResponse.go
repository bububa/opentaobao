package caipiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoGoodsInfoInputAPIResponse 录入参加送彩票商品信息 API返回值
// taobao.caipiao.goods.info.input
//
// 录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。
type TaobaoCaipiaoGoodsInfoInputAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoGoodsInfoInputAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCaipiaoGoodsInfoInputAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCaipiaoGoodsInfoInputAPIResponseModel).Reset()
}

// TaobaoCaipiaoGoodsInfoInputAPIResponseModel is 录入参加送彩票商品信息 成功返回结果
type TaobaoCaipiaoGoodsInfoInputAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_goods_info_input_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 录入操作是否成功
	InputResult bool `json:"input_result,omitempty" xml:"input_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCaipiaoGoodsInfoInputAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InputResult = false
}

var poolTaobaoCaipiaoGoodsInfoInputAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCaipiaoGoodsInfoInputAPIResponse)
	},
}

// GetTaobaoCaipiaoGoodsInfoInputAPIResponse 从 sync.Pool 获取 TaobaoCaipiaoGoodsInfoInputAPIResponse
func GetTaobaoCaipiaoGoodsInfoInputAPIResponse() *TaobaoCaipiaoGoodsInfoInputAPIResponse {
	return poolTaobaoCaipiaoGoodsInfoInputAPIResponse.Get().(*TaobaoCaipiaoGoodsInfoInputAPIResponse)
}

// ReleaseTaobaoCaipiaoGoodsInfoInputAPIResponse 将 TaobaoCaipiaoGoodsInfoInputAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCaipiaoGoodsInfoInputAPIResponse(v *TaobaoCaipiaoGoodsInfoInputAPIResponse) {
	v.Reset()
	poolTaobaoCaipiaoGoodsInfoInputAPIResponse.Put(v)
}
