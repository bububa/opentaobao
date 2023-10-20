package moscm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsSynchinventorybycountingAPIResponse 以盘点方式调整库存：传入商品实际库存 API返回值
// alibaba.mos.goods.synchinventorybycounting
//
// 以盘点方式调整库存：传入商品实际库存
// 盘点单自动判断数量增减
type AlibabaMosGoodsSynchinventorybycountingAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsSynchinventorybycountingAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSynchinventorybycountingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosGoodsSynchinventorybycountingAPIResponseModel).Reset()
}

// AlibabaMosGoodsSynchinventorybycountingAPIResponseModel is 以盘点方式调整库存：传入商品实际库存 成功返回结果
type AlibabaMosGoodsSynchinventorybycountingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_synchinventorybycounting_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回盘点单号
	Result *AlibabaMosGoodsSynchinventorybycountingResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosGoodsSynchinventorybycountingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosGoodsSynchinventorybycountingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosGoodsSynchinventorybycountingAPIResponse)
	},
}

// GetAlibabaMosGoodsSynchinventorybycountingAPIResponse 从 sync.Pool 获取 AlibabaMosGoodsSynchinventorybycountingAPIResponse
func GetAlibabaMosGoodsSynchinventorybycountingAPIResponse() *AlibabaMosGoodsSynchinventorybycountingAPIResponse {
	return poolAlibabaMosGoodsSynchinventorybycountingAPIResponse.Get().(*AlibabaMosGoodsSynchinventorybycountingAPIResponse)
}

// ReleaseAlibabaMosGoodsSynchinventorybycountingAPIResponse 将 AlibabaMosGoodsSynchinventorybycountingAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosGoodsSynchinventorybycountingAPIResponse(v *AlibabaMosGoodsSynchinventorybycountingAPIResponse) {
	v.Reset()
	poolAlibabaMosGoodsSynchinventorybycountingAPIResponse.Put(v)
}
