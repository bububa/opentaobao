package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosIsvInventoryScrollqueryAPIResponse 滚动查询库存数据 API返回值
// alibaba.mos.isv.inventory.scrollquery
//
// 按专柜滚动查询有库存商品
type AlibabaMosIsvInventoryScrollqueryAPIResponse struct {
	model.CommonResponse
	AlibabaMosIsvInventoryScrollqueryAPIResponseModel
}

// AlibabaMosIsvInventoryScrollqueryAPIResponseModel is 滚动查询库存数据 成功返回结果
type AlibabaMosIsvInventoryScrollqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_isv_inventory_scrollquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MosScrollQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
