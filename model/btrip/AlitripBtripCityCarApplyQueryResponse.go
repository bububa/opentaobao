package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方市内用车申请单查询 API返回值 
alitrip.btrip.city.car.apply.query

三方市内用车申请单查询
*/
type AlitripBtripCityCarApplyQueryAPIResponse struct {
    model.CommonResponse
    AlitripBtripCityCarApplyQueryResponse
}

// 三方市内用车申请单查询 成功返回结果
type AlitripBtripCityCarApplyQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_city_car_apply_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
