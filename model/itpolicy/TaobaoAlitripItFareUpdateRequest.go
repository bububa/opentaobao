package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条修改 APIRequest
taobao.alitrip.it.fare.update

自有政策修改接口，可以根据fareId或outId修改，outId不唯一时，不能用outId修改。当外部政策id、出发城市、到达城市、出票航司任一有变化，或往返时是否允许混舱、文件编号、可混文件编号任一有变化，将删除老数据，产生一条新政策。
*/
type TaobaoAlitripItFareUpdateRequest struct {
    model.Params

    // json格式的字符串，扩展属性，预留
    extendAttributes   string 

    // 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
    fareId   int64 

    // 自有政策json序列化字符串，具体属性视fareType的类型，和addow或addrt接口的请求参数一致，如：{"ticketingAirline":"CA","saleAirline":"CA","addressOption":"城市","tripType":"直达","originLand":"SHA,BJS","destination":"HKG,SHA","cabin":"N","restrictFlightNo":"","excludeFlightNo":"","validDate4Dep":"2015-05-30~2015-11-30","flightDateRestrict4Dep":"12","saleDate":"2015-05-30~2015-11-30","adultPassengerIdentity":"普通","ticketPrice":2000,"childPrice":"1980","returnPoint":1.0,"adjustMoney":0,"refundRule":"收取80000元退票费","reissueRule":"收取20%改期费","noshowRule":"起飞前不得退票，不得改期","luggageRule":"逾重行李费用为每公斤100元"}
    fareJson   string 

    // 运价类型，1单程 2往返
    fareType   int64 

    // 外部id，为新增时请求参数中的外部政策id
    outId   string 

}

func NewTaobaoAlitripItFareUpdateRequest() *TaobaoAlitripItFareUpdateRequest{
    return &TaobaoAlitripItFareUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItFareUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.update"
}

func (r TaobaoAlitripItFareUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItFareUpdateRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItFareUpdateRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItFareUpdateRequest) SetFareId(fareId int64) error {
    r.fareId = fareId
    r.Set("fareId", fareId)
    return nil
}

func (r TaobaoAlitripItFareUpdateRequest) GetFareId() int64 {
    return r.fareId
}

func (r *TaobaoAlitripItFareUpdateRequest) SetFareJson(fareJson string) error {
    r.fareJson = fareJson
    r.Set("fareJson", fareJson)
    return nil
}

func (r TaobaoAlitripItFareUpdateRequest) GetFareJson() string {
    return r.fareJson
}

func (r *TaobaoAlitripItFareUpdateRequest) SetFareType(fareType int64) error {
    r.fareType = fareType
    r.Set("fareType", fareType)
    return nil
}

func (r TaobaoAlitripItFareUpdateRequest) GetFareType() int64 {
    return r.fareType
}

func (r *TaobaoAlitripItFareUpdateRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("outId", outId)
    return nil
}

func (r TaobaoAlitripItFareUpdateRequest) GetOutId() string {
    return r.outId
}

