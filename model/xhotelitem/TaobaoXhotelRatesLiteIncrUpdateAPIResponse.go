package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店价格库存轻量级增量接口 API返回值 
taobao.xhotel.rates.lite.incr.update

多个rate的库存房价开关的增量更新接口
*/
type TaobaoXhotelRatesLiteIncrUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel
}

// 酒店价格库存轻量级增量接口 成功返回结果
type TaobaoXhotelRatesLiteIncrUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_rates_lite_incr_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoXhotelRatesLiteIncrUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
