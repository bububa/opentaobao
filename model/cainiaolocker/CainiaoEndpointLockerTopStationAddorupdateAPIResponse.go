package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加更新代收点 API返回值 
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息
*/
type CainiaoEndpointLockerTopStationAddorupdateAPIResponse struct {
    model.CommonResponse
    CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel
}

// 增加更新代收点 成功返回结果
type CainiaoEndpointLockerTopStationAddorupdateAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_endpoint_locker_top_station_addorupdate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
