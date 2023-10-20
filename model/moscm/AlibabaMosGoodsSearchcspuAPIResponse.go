package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSearchcspuAPIResponse cspu查询 API返回值
// alibaba.mos.goods.searchcspu
//
// 商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
type AlibabaMosGoodsSearchcspuAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsSearchcspuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSearchcspuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsSearchcspuAPIResponseModel).Reset()
}

// AlibabaMosGoodsSearchcspuAPIResponseModel is cspu查询 成功返回结果
type AlibabaMosGoodsSearchcspuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_searchcspu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *PagedList `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSearchcspuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaMosGoodsSearchcspuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsSearchcspuAPIResponse)
	},
}

// GetAlibabaMosGoodsSearchcspuAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsSearchcspuAPIResponse
func GetAlibabaMosGoodsSearchcspuAPIResponse() *AlibabaMosGoodsSearchcspuAPIResponse {
	return poolAlibabaMosGoodsSearchcspuAPIResponse.Get().(*AlibabaMosGoodsSearchcspuAPIResponse)
}

// ReleaseAlibabaMosGoodsSearchcspuAPIResponse 将 AlibabaMosGoodsSearchcspuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsSearchcspuAPIResponse(v *AlibabaMosGoodsSearchcspuAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsSearchcspuAPIResponse.Put(v)
}
