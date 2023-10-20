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
	// 自有政策json序列化后字符串，具体属性和addow或addrt接口的请求参数一致，如：{&#34;ticketingAirline&#34;:&#34;CA&#34;,&#34;saleAirline&#34;:&#34;CA&#34;,&#34;addressOption&#34;:&#34;城市&#34;,&#34;tripType&#34;:&#34;直达&#34;,&#34;originLand&#34;:&#34;SHA,BJS&#34;,&#34;destination&#34;:&#34;HKG,SHA&#34;,&#34;cabin&#34;:&#34;N&#34;,&#34;restrictFlightNo&#34;:&#34;&#34;,&#34;excludeFlightNo&#34;:&#34;&#34;,&#34;validDate4Dep&#34;:&#34;2015-05-30~2015-11-30&#34;,&#34;flightDateRestrict4Dep&#34;:&#34;12&#34;,&#34;saleDate&#34;:&#34;2015-05-30~2015-11-30&#34;,&#34;adultPassengerIdentity&#34;:&#34;普通&#34;,&#34;ticketPrice&#34;:2000,&#34;childPrice&#34;:&#34;1980&#34;,&#34;returnPoint&#34;:1.0,&#34;adjustMoney&#34;:0,&#34;refundRule&#34;:&#34;收取80000元退票费&#34;,&#34;reissueRule&#34;:&#34;收取20%改期费&#34;,&#34;noshowRule&#34;:&#34;起飞前不得退票，不得改期&#34;,&#34;luggageRule&#34;:&#34;逾重行李费用为每公斤100元&#34;}
	FareJson string `json:"fare_json,omitempty" xml:"fare_json,omitempty"`
	// 运价id
	FareId int64 `json:"fare_id,omitempty" xml:"fare_id,omitempty"`
	// 运价类型，1单程 2往返
	FareType int64 `json:"fare_type,omitempty" xml:"fare_type,omitempty"`
	// 0：未发布 1：已发布 2：已过期
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
