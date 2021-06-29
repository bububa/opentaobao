package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条删除 APIRequest
taobao.alitrip.it.fare.delete

自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
*/
type TaobaoAlitripItFareDeleteRequest struct {
    model.Params

    // json格式的字符串，扩展属性，预留
    extendAttributes   string 

    // 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
    fareId   int64 

    // 外部id，为新增时请求参数中的外部政策id
    outId   string 

}

func NewTaobaoAlitripItFareDeleteRequest() *TaobaoAlitripItFareDeleteRequest{
    return &TaobaoAlitripItFareDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripItFareDeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.delete"
}

func (r TaobaoAlitripItFareDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripItFareDeleteRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

func (r TaobaoAlitripItFareDeleteRequest) GetExtendAttributes() string {
    return r.extendAttributes
}

func (r *TaobaoAlitripItFareDeleteRequest) SetFareId(fareId int64) error {
    r.fareId = fareId
    r.Set("fareId", fareId)
    return nil
}

func (r TaobaoAlitripItFareDeleteRequest) GetFareId() int64 {
    return r.fareId
}

func (r *TaobaoAlitripItFareDeleteRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("outId", outId)
    return nil
}

func (r TaobaoAlitripItFareDeleteRequest) GetOutId() string {
    return r.outId
}

