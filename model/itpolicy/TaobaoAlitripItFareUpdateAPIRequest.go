package itpolicy

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripItFareUpdateAPIRequest
【国际机票自有政策】单条修改 API请求
taobao.alitrip.it.fare.update

自有政策修改接口，可以根据fareId或outId修改，outId不唯一时，不能用outId修改。当外部政策id、出发城市、到达城市、出票航司任一有变化，或往返时是否允许混舱、文件编号、可混文件编号任一有变化，将删除老数据，产生一条新政策。 */
type TaobaoAlitripItFareUpdateAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
	_fareId int64
	// 自有政策json序列化字符串，具体属性视fareType的类型，和addow或addrt接口的请求参数一致，如：{"ticketingAirline":"CA","saleAirline":"CA","addressOption":"城市","tripType":"直达","originLand":"SHA,BJS","destination":"HKG,SHA","cabin":"N","restrictFlightNo":"","excludeFlightNo":"","validDate4Dep":"2015-05-30~2015-11-30","flightDateRestrict4Dep":"12","saleDate":"2015-05-30~2015-11-30","adultPassengerIdentity":"普通","ticketPrice":2000,"childPrice":"1980","returnPoint":1.0,"adjustMoney":0,"refundRule":"收取80000元退票费","reissueRule":"收取20%改期费","noshowRule":"起飞前不得退票，不得改期","luggageRule":"逾重行李费用为每公斤100元"}
	_fareJson string
	// 运价类型，1单程 2往返
	_fareType int64
	// 外部id，为新增时请求参数中的外部政策id
	_outId string
}

// NewTaobaoAlitripItFareUpdateRequest 初始化TaobaoAlitripItFareUpdateAPIRequest对象
func NewTaobaoAlitripItFareUpdateRequest() *TaobaoAlitripItFareUpdateAPIRequest {
	return &TaobaoAlitripItFareUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareUpdateAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// Get ExtendAttributes Getter
func (r TaobaoAlitripItFareUpdateAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// Set is FareId Setter
// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
func (r *TaobaoAlitripItFareUpdateAPIRequest) SetFareId(_fareId int64) error {
	r._fareId = _fareId
	r.Set("fareId", _fareId)
	return nil
}

// Get FareId Getter
func (r TaobaoAlitripItFareUpdateAPIRequest) GetFareId() int64 {
	return r._fareId
}

// Set is FareJson Setter
// 自有政策json序列化字符串，具体属性视fareType的类型，和addow或addrt接口的请求参数一致，如：{"ticketingAirline":"CA","saleAirline":"CA","addressOption":"城市","tripType":"直达","originLand":"SHA,BJS","destination":"HKG,SHA","cabin":"N","restrictFlightNo":"","excludeFlightNo":"","validDate4Dep":"2015-05-30~2015-11-30","flightDateRestrict4Dep":"12","saleDate":"2015-05-30~2015-11-30","adultPassengerIdentity":"普通","ticketPrice":2000,"childPrice":"1980","returnPoint":1.0,"adjustMoney":0,"refundRule":"收取80000元退票费","reissueRule":"收取20%改期费","noshowRule":"起飞前不得退票，不得改期","luggageRule":"逾重行李费用为每公斤100元"}
func (r *TaobaoAlitripItFareUpdateAPIRequest) SetFareJson(_fareJson string) error {
	r._fareJson = _fareJson
	r.Set("fareJson", _fareJson)
	return nil
}

// Get FareJson Getter
func (r TaobaoAlitripItFareUpdateAPIRequest) GetFareJson() string {
	return r._fareJson
}

// Set is FareType Setter
// 运价类型，1单程 2往返
func (r *TaobaoAlitripItFareUpdateAPIRequest) SetFareType(_fareType int64) error {
	r._fareType = _fareType
	r.Set("fareType", _fareType)
	return nil
}

// Get FareType Getter
func (r TaobaoAlitripItFareUpdateAPIRequest) GetFareType() int64 {
	return r._fareType
}

// Set is OutId Setter
// 外部id，为新增时请求参数中的外部政策id
func (r *TaobaoAlitripItFareUpdateAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("outId", _outId)
	return nil
}

// Get OutId Getter
func (r TaobaoAlitripItFareUpdateAPIRequest) GetOutId() string {
	return r._outId
}
