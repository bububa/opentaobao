package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightbasicdatacityqueryAllAPIResponse 机票基础数据城市数据查询 API返回值
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
type AlitripflightbasicdatacityqueryAllAPIResponse struct {
	model.CommonResponse
	AlitripflightbasicdatacityqueryAllAPIResponseModel
}

// AlitripflightbasicdatacityqueryAllAPIResponseModel is 机票基础数据城市数据查询 成功返回结果
type AlitripflightbasicdatacityqueryAllAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_basic_data_city_queryAll_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
