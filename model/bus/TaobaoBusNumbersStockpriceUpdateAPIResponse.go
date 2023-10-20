package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusnumbersstockpriceupdateAPIResponse 汽车票更新价格库存 API返回值
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
type TaobaobusnumbersstockpriceupdateAPIResponse struct {
	model.CommonResponse
	TaobaobusnumbersstockpriceupdateAPIResponseModel
}

// TaobaobusnumbersstockpriceupdateAPIResponseModel is 汽车票更新价格库存 成功返回结果
type TaobaobusnumbersstockpriceupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_numbers_stockprice_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功数量
	SuccCount int64 `json:"succ_count,omitempty" xml:"succ_count,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
