package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
终端配置支持 API请求
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回
*/
type TaobaoLarkPosItemprodFindterminalRequest struct {
    model.Params
    // 终端id
    deviceId   string
    // 终端类型
    deviceType   string
    // 912874323429834
    createUser   string
    // 租户编码
    leaseCode   string
    // 影城id
    cinemaId   string
    // 影城名称
    cinemaName   string
}

// 初始化TaobaoLarkPosItemprodFindterminalRequest对象
func NewTaobaoLarkPosItemprodFindterminalRequest() *TaobaoLarkPosItemprodFindterminalRequest{
    return &TaobaoLarkPosItemprodFindterminalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkPosItemprodFindterminalRequest) GetApiMethodName() string {
    return "taobao.lark.pos.itemprod.findterminal"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkPosItemprodFindterminalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 终端id
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetDeviceId() string {
    return r.deviceId
}
// DeviceType Setter
// 终端类型
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetDeviceType() string {
    return r.deviceType
}
// CreateUser Setter
// 912874323429834
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetCreateUser(createUser string) error {
    r.createUser = createUser
    r.Set("create_user", createUser)
    return nil
}

// CreateUser Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetCreateUser() string {
    return r.createUser
}
// LeaseCode Setter
// 租户编码
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetLeaseCode(leaseCode string) error {
    r.leaseCode = leaseCode
    r.Set("lease_code", leaseCode)
    return nil
}

// LeaseCode Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetLeaseCode() string {
    return r.leaseCode
}
// CinemaId Setter
// 影城id
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetCinemaId(cinemaId string) error {
    r.cinemaId = cinemaId
    r.Set("cinema_id", cinemaId)
    return nil
}

// CinemaId Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetCinemaId() string {
    return r.cinemaId
}
// CinemaName Setter
// 影城名称
func (r *TaobaoLarkPosItemprodFindterminalRequest) SetCinemaName(cinemaName string) error {
    r.cinemaName = cinemaName
    r.Set("cinema_name", cinemaName)
    return nil
}

// CinemaName Getter
func (r TaobaoLarkPosItemprodFindterminalRequest) GetCinemaName() string {
    return r.cinemaName
}
