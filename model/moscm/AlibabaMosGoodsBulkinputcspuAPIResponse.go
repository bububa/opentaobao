package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsBulkinputcspuAPIResponse 批量录入商品信息 API返回值
// alibaba.mos.goods.bulkinputcspu
//
// 用于商品信息的批量导入到银泰商品中台
type AlibabaMosGoodsBulkinputcspuAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsBulkinputcspuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsBulkinputcspuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsBulkinputcspuAPIResponseModel).Reset()
}

// AlibabaMosGoodsBulkinputcspuAPIResponseModel is 批量录入商品信息 成功返回结果
type AlibabaMosGoodsBulkinputcspuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_bulkinputcspu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Data *BulkInputCspuResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsBulkinputcspuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolAlibabaMosGoodsBulkinputcspuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsBulkinputcspuAPIResponse)
	},
}

// GetAlibabaMosGoodsBulkinputcspuAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsBulkinputcspuAPIResponse
func GetAlibabaMosGoodsBulkinputcspuAPIResponse() *AlibabaMosGoodsBulkinputcspuAPIResponse {
	return poolAlibabaMosGoodsBulkinputcspuAPIResponse.Get().(*AlibabaMosGoodsBulkinputcspuAPIResponse)
}

// ReleaseAlibabaMosGoodsBulkinputcspuAPIResponse 将 AlibabaMosGoodsBulkinputcspuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsBulkinputcspuAPIResponse(v *AlibabaMosGoodsBulkinputcspuAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsBulkinputcspuAPIResponse.Put(v)
}
