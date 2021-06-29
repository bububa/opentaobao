package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交互卡片预约信息读取接口 API请求
taobao.oc.eservice.appoint.list

允许外部的isv通过该接口读取门店预约信息
*/
type TaobaoOcEserviceAppointListRequest struct {
    model.Params
    // 预约信息唯一编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
    _code   int64
    // 门店编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
    _mallCode   string
    // 查询预约的起始时间，格式yyyyMMddHHmmss，默认为当前时间
    _startAppointTime   string
    // 买家客户nick(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
    _customerNick   string
    // 买家客户电话号码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
    _customerPhone   string
    // 买家客户装修房屋所在的市(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
    _houseAddressCity   string
    // 卖家主账号id
    _sellerId   int64
    // 返回结果按预约时间排序，指示升序还是降息，取值asc和desc
    _sortOrder   string
}

// 初始化TaobaoOcEserviceAppointListRequest对象
func NewTaobaoOcEserviceAppointListRequest() *TaobaoOcEserviceAppointListRequest{
    return &TaobaoOcEserviceAppointListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcEserviceAppointListRequest) GetApiMethodName() string {
    return "taobao.oc.eservice.appoint.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcEserviceAppointListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 预约信息唯一编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListRequest) SetCode(_code int64) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r TaobaoOcEserviceAppointListRequest) GetCode() int64 {
    return r._code
}
// MallCode Setter
// 门店编码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListRequest) SetMallCode(_mallCode string) error {
    r._mallCode = _mallCode
    r.Set("mall_code", _mallCode)
    return nil
}

// MallCode Getter
func (r TaobaoOcEserviceAppointListRequest) GetMallCode() string {
    return r._mallCode
}
// StartAppointTime Setter
// 查询预约的起始时间，格式yyyyMMddHHmmss，默认为当前时间
func (r *TaobaoOcEserviceAppointListRequest) SetStartAppointTime(_startAppointTime string) error {
    r._startAppointTime = _startAppointTime
    r.Set("start_appoint_time", _startAppointTime)
    return nil
}

// StartAppointTime Getter
func (r TaobaoOcEserviceAppointListRequest) GetStartAppointTime() string {
    return r._startAppointTime
}
// CustomerNick Setter
// 买家客户nick(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListRequest) SetCustomerNick(_customerNick string) error {
    r._customerNick = _customerNick
    r.Set("customer_nick", _customerNick)
    return nil
}

// CustomerNick Getter
func (r TaobaoOcEserviceAppointListRequest) GetCustomerNick() string {
    return r._customerNick
}
// CustomerPhone Setter
// 买家客户电话号码(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListRequest) SetCustomerPhone(_customerPhone string) error {
    r._customerPhone = _customerPhone
    r.Set("customer_phone", _customerPhone)
    return nil
}

// CustomerPhone Getter
func (r TaobaoOcEserviceAppointListRequest) GetCustomerPhone() string {
    return r._customerPhone
}
// HouseAddressCity Setter
// 买家客户装修房屋所在的市(code, customerNick, customerPhone, houseAddressCity, mallCode 调用时五个可选参数中任选一个作为输入参数)
func (r *TaobaoOcEserviceAppointListRequest) SetHouseAddressCity(_houseAddressCity string) error {
    r._houseAddressCity = _houseAddressCity
    r.Set("house_address_city", _houseAddressCity)
    return nil
}

// HouseAddressCity Getter
func (r TaobaoOcEserviceAppointListRequest) GetHouseAddressCity() string {
    return r._houseAddressCity
}
// SellerId Setter
// 卖家主账号id
func (r *TaobaoOcEserviceAppointListRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoOcEserviceAppointListRequest) GetSellerId() int64 {
    return r._sellerId
}
// SortOrder Setter
// 返回结果按预约时间排序，指示升序还是降息，取值asc和desc
func (r *TaobaoOcEserviceAppointListRequest) SetSortOrder(_sortOrder string) error {
    r._sortOrder = _sortOrder
    r.Set("sort_order", _sortOrder)
    return nil
}

// SortOrder Getter
func (r TaobaoOcEserviceAppointListRequest) GetSortOrder() string {
    return r._sortOrder
}
