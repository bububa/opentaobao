package itpolicy

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票自有政策】单条删除 API请求
taobao.alitrip.it.fare.delete

自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
*/
type TaobaoAlitripItFareDeleteAPIRequest struct {
    model.Params
    // json格式的字符串，扩展属性，预留
    _extendAttributes   string
    // 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
    _fareId   int64
    // 外部id，为新增时请求参数中的外部政策id
    _outId   string
}

// 初始化TaobaoAlitripItFareDeleteAPIRequest对象
func NewTaobaoAlitripItFareDeleteRequest() *TaobaoAlitripItFareDeleteAPIRequest{
    return &TaobaoAlitripItFareDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.it.fare.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetExtendAttributes(_extendAttributes string) error {
    r._extendAttributes = _extendAttributes
    r.Set("extendAttributes", _extendAttributes)
    return nil
}

// ExtendAttributes Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetExtendAttributes() string {
    return r._extendAttributes
}
// FareId Setter
// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetFareId(_fareId int64) error {
    r._fareId = _fareId
    r.Set("fareId", _fareId)
    return nil
}

// FareId Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetFareId() int64 {
    return r._fareId
}
// OutId Setter
// 外部id，为新增时请求参数中的外部政策id
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("outId", _outId)
    return nil
}

// OutId Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetOutId() string {
    return r._outId
}
