package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条查询 API请求
taobao.alitrip.it.fare.get

通过此接口可以查询单条政策的详情，可以根据fareId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据
*/
type TaobaoAlitripItFareGetRequest struct {
    model.Params
    // json格式的字符串，扩展属性，预留
    extendAttributes   string
    // 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
    fareId   int64
    // 外部id，为新增时请求参数中的外部政策id
    outId   string
}

// 初始化TaobaoAlitripItFareGetRequest对象
func NewTaobaoAlitripItFareGetRequest() *TaobaoAlitripItFareGetRequest{
    return &TaobaoAlitripItFareGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareGetRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareGetRequest) SetExtendAttributes(extendAttributes string) error {
    r.extendAttributes = extendAttributes
    r.Set("extendAttributes", extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareGetRequest) GetExtendAttributes() string {
    return r.extendAttributes
}
// FareId Setter
// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
func (r *TaobaoAlitripItFareGetRequest) SetFareId(fareId int64) error {
    r.fareId = fareId
    r.Set("fareId", fareId)
    return nil
}

// FareId Getter
func (r TaobaoAlitripItFareGetRequest) GetFareId() int64 {
    return r.fareId
}
// OutId Setter
// 外部id，为新增时请求参数中的外部政策id
func (r *TaobaoAlitripItFareGetRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("outId", outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItFareGetRequest) GetOutId() string {
    return r.outId
}
