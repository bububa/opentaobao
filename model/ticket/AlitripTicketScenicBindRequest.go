package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票景点绑定接口 API请求
alitrip.ticket.scenic.bind

门票景点绑定接口，用于建立阿里标准景点id与商家系统景点id的映射关系。该接口同时支持新建和修改映射关系，当用户没有为ali_scenic_id建立过映射关系时，则判断为新建映射关系，否则为修改。可以通过设置update_out_scenic_id来修改ali_scenic_id与out_scenic_id的映射关系。
*/
type AlitripTicketScenicBindRequest struct {
    model.Params
    // 必填，阿里旅行对应的景点编码
    aliScenicId   int64
    // 商户景点城市
    city   string
    // 商户景点地址
    address   string
    // 商户景点名称
    outScenicName   string
    // 商户景点省份
    province   string
    // 必填，商户系统中的景点编码，用于与ali_scenic_id建立映射关系
    outScenicId   string
    // 可选，如果需要更新out_scenic_id与ali_scenic_id的映射关系时 需要填写
    updateOutScenicId   string
}

// 初始化AlitripTicketScenicBindRequest对象
func NewAlitripTicketScenicBindRequest() *AlitripTicketScenicBindRequest{
    return &AlitripTicketScenicBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTicketScenicBindRequest) GetApiMethodName() string {
    return "alitrip.ticket.scenic.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTicketScenicBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliScenicId Setter
// 必填，阿里旅行对应的景点编码
func (r *AlitripTicketScenicBindRequest) SetAliScenicId(aliScenicId int64) error {
    r.aliScenicId = aliScenicId
    r.Set("ali_scenic_id", aliScenicId)
    return nil
}

// AliScenicId Getter
func (r AlitripTicketScenicBindRequest) GetAliScenicId() int64 {
    return r.aliScenicId
}
// City Setter
// 商户景点城市
func (r *AlitripTicketScenicBindRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlitripTicketScenicBindRequest) GetCity() string {
    return r.city
}
// Address Setter
// 商户景点地址
func (r *AlitripTicketScenicBindRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlitripTicketScenicBindRequest) GetAddress() string {
    return r.address
}
// OutScenicName Setter
// 商户景点名称
func (r *AlitripTicketScenicBindRequest) SetOutScenicName(outScenicName string) error {
    r.outScenicName = outScenicName
    r.Set("out_scenic_name", outScenicName)
    return nil
}

// OutScenicName Getter
func (r AlitripTicketScenicBindRequest) GetOutScenicName() string {
    return r.outScenicName
}
// Province Setter
// 商户景点省份
func (r *AlitripTicketScenicBindRequest) SetProvince(province string) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r AlitripTicketScenicBindRequest) GetProvince() string {
    return r.province
}
// OutScenicId Setter
// 必填，商户系统中的景点编码，用于与ali_scenic_id建立映射关系
func (r *AlitripTicketScenicBindRequest) SetOutScenicId(outScenicId string) error {
    r.outScenicId = outScenicId
    r.Set("out_scenic_id", outScenicId)
    return nil
}

// OutScenicId Getter
func (r AlitripTicketScenicBindRequest) GetOutScenicId() string {
    return r.outScenicId
}
// UpdateOutScenicId Setter
// 可选，如果需要更新out_scenic_id与ali_scenic_id的映射关系时 需要填写
func (r *AlitripTicketScenicBindRequest) SetUpdateOutScenicId(updateOutScenicId string) error {
    r.updateOutScenicId = updateOutScenicId
    r.Set("update_out_scenic_id", updateOutScenicId)
    return nil
}

// UpdateOutScenicId Getter
func (r AlitripTicketScenicBindRequest) GetUpdateOutScenicId() string {
    return r.updateOutScenicId
}
