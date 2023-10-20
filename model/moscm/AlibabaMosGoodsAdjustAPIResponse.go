package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsAdjustAPIResponse 调整库存 API返回值
// alibaba.mos.goods.adjust
//
// 库存调整接口
type AlibabaMosGoodsAdjustAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsAdjustAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsAdjustAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsAdjustAPIResponseModel).Reset()
}

// AlibabaMosGoodsAdjustAPIResponseModel is 调整库存 成功返回结果
type AlibabaMosGoodsAdjustAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存调整单号
	Result *AlibabaMosGoodsAdjustResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsAdjustAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosGoodsAdjustAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsAdjustAPIResponse)
	},
}

// GetAlibabaMosGoodsAdjustAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsAdjustAPIResponse
func GetAlibabaMosGoodsAdjustAPIResponse() *AlibabaMosGoodsAdjustAPIResponse {
	return poolAlibabaMosGoodsAdjustAPIResponse.Get().(*AlibabaMosGoodsAdjustAPIResponse)
}

// ReleaseAlibabaMosGoodsAdjustAPIResponse 将 AlibabaMosGoodsAdjustAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsAdjustAPIResponse(v *AlibabaMosGoodsAdjustAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsAdjustAPIResponse.Put(v)
}
