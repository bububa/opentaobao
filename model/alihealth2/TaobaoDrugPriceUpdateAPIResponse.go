package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新宝贝价格 API返回值 
taobao.drug.price.update

商家更新价格
*/
type TaobaoDrugPriceUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoDrugPriceUpdateAPIResponseModel
}

// 商家更新宝贝价格 成功返回结果
type TaobaoDrugPriceUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"drug_price_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoDrugPriceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
