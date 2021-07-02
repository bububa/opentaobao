package itpolicy

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareGetAPIResponse 【国际机票自有政策】单条查询 API返回值
// taobao.alitrip.it.fare.get
//
// 通过此接口可以查询单条政策的详情，可以根据fareId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据
type TaobaoAlitripItFareGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripItFareGetAPIResponseModel
}

// TaobaoAlitripItFareGetAPIResponseModel is 【国际机票自有政策】单条查询 成功返回结果
type TaobaoAlitripItFareGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_it_fare_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// json格式的字符串，扩展属性，预留
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 运价id
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
	// 自有政策json序列化后字符串，具体属性和addow或addrt接口的请求参数一致，如：{"ticketingAirline":"CA","saleAirline":"CA","addressOption":"城市","tripType":"直达","originLand":"SHA,BJS","destination":"HKG,SHA","cabin":"N","restrictFlightNo":"","excludeFlightNo":"","validDate4Dep":"2015-05-30~2015-11-30","flightDateRestrict4Dep":"12","saleDate":"2015-05-30~2015-11-30","adultPassengerIdentity":"普通","ticketPrice":2000,"childPrice":"1980","returnPoint":1.0,"adjustMoney":0,"refundRule":"收取80000元退票费","reissueRule":"收取20%改期费","noshowRule":"起飞前不得退票，不得改期","luggageRule":"逾重行李费用为每公斤100元"}
	FareJson string `json:"fare_json,omitempty" xml:"fare_json,omitempty"`
	// 运价类型，1单程 2往返
	FareType int64 `json:"fare_type,omitempty" xml:"fare_type,omitempty"`
	// 0：未发布 1：已发布 2：已过期
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
