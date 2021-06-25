package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素管理接口 APIRequest
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
*/
type TaobaoAlitripTravelItemElementManageRequest struct {
    model.Params

    // 必填，操作类型：1-新增，2-修改，3-删除。。特别注意：删除 为逻辑删除，即该outer_id所对应的元素还存在但是会置为无效状态，重新编辑修改即可恢复为有效状态。因此该id一旦使用将不可重复
    operation   int64 

    // 必填，元素的外部商家编码，必须唯一。编辑、删除时将根据该编码找到对应元素。
    outerId   string 

    // 资源元素类型，新增时必填：1-景点，2-酒店，5-交通接驳，6-WIFI库，7-电话卡，8-餐饮，9-签证库，11-特色活动，999-其他
    elementType   int64 

    // 元素名称，新增时必填； 注意：Wifi库的使用地和签证库所在国家均适用这个字段
    name   string 

    // 元素所在城市，景点、酒店在新增时必填
    city   string 

    // 元素的子类型，新增时必填。景点指门票类型，酒店指房型信息，交通接驳（接送机、接驳车、租车、船票、其他）选其一，餐饮（早餐、晚餐、午餐、下午茶及其他）选其一；签证（旅游签证、商务签证、工作签证、留学签证、探亲访友签证、入台证、其他）
    type   string 

    // 当新增“交通接驳、餐饮、特色活动、其他”资源类型时 必填
    desc   string 

}

func NewTaobaoAlitripTravelItemElementManageRequest() *TaobaoAlitripTravelItemElementManageRequest{
    return &TaobaoAlitripTravelItemElementManageRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.element.manage"
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelItemElementManageRequest) SetOperation(operation int64) error {
    r.operation = operation
    r.Set("operation", operation)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetOperation() int64 {
    return r.operation
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetElementType(elementType int64) error {
    r.elementType = elementType
    r.Set("element_type", elementType)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetElementType() int64 {
    return r.elementType
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetName() string {
    return r.name
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetCity() string {
    return r.city
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetType() string {
    return r.type
}

func (r *TaobaoAlitripTravelItemElementManageRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoAlitripTravelItemElementManageRequest) GetDesc() string {
    return r.desc
}

