package moscm

import (
	"encoding/xml"

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

// AlibabaMosGoodsSynchinventorybycountingAPIResponseModel is 以盘点方式调整库存：传入商品实际库存 成功返回结果
type AlibabaMosGoodsSynchinventorybycountingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_synchinventorybycounting_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回盘点单号
	Result *AlibabaMosGoodsSynchinventorybycountingResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
