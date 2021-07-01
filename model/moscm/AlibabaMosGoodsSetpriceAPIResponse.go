package moscm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
价格变更接口 API返回值 
alibaba.mos.goods.setprice

价格变更接口，供供应商修改价格时使用。
*/
type AlibabaMosGoodsSetpriceAPIResponse struct {
    model.CommonResponse
    AlibabaMosGoodsSetpriceAPIResponseModel
}

// 价格变更接口 成功返回结果
type AlibabaMosGoodsSetpriceAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mos_goods_setprice_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    DataS   []PriceResult `json:"data_s,omitempty" xml:"data_s>price_result,omitempty"`
}
