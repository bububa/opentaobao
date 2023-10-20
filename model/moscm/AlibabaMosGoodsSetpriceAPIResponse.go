package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSetpriceAPIResponse 价格变更接口 API返回值
// alibaba.mos.goods.setprice
//
// 价格变更接口，供供应商修改价格时使用。
type AlibabaMosGoodsSetpriceAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsSetpriceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSetpriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsSetpriceAPIResponseModel).Reset()
}

// AlibabaMosGoodsSetpriceAPIResponseModel is 价格变更接口 成功返回结果
type AlibabaMosGoodsSetpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_setprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	DataS []PriceResult `json:"data_s,omitempty" xml:"data_s>price_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSetpriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataS = m.DataS[:0]
}

var poolAlibabaMosGoodsSetpriceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsSetpriceAPIResponse)
	},
}

// GetAlibabaMosGoodsSetpriceAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsSetpriceAPIResponse
func GetAlibabaMosGoodsSetpriceAPIResponse() *AlibabaMosGoodsSetpriceAPIResponse {
	return poolAlibabaMosGoodsSetpriceAPIResponse.Get().(*AlibabaMosGoodsSetpriceAPIResponse)
}

// ReleaseAlibabaMosGoodsSetpriceAPIResponse 将 AlibabaMosGoodsSetpriceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsSetpriceAPIResponse(v *AlibabaMosGoodsSetpriceAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsSetpriceAPIResponse.Put(v)
}
