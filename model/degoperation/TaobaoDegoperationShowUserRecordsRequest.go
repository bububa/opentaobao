package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户中奖记录 API请求
taobao.degoperation.show.user.records

用户中奖记录
*/
type TaobaoDegoperationShowUserRecordsRequest struct {
    model.Params
    // 活动后台配置
    _degAppKey   string
    // 活动后台配置
    _eventKey   string
    // 第几页
    _pageNumber   int64
    // 分页尺寸
    _pageSize   int64
    // 系统信息
    _degAccessToken   string
}

// 初始化TaobaoDegoperationShowUserRecordsRequest对象
func NewTaobaoDegoperationShowUserRecordsRequest() *TaobaoDegoperationShowUserRecordsRequest{
    return &TaobaoDegoperationShowUserRecordsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationShowUserRecordsRequest) GetApiMethodName() string {
    return "taobao.degoperation.show.user.records"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationShowUserRecordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DegAppKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAppKey(_degAppKey string) error {
    r._degAppKey = _degAppKey
    r.Set("deg_app_key", _degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAppKey() string {
    return r._degAppKey
}
// EventKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowUserRecordsRequest) SetEventKey(_eventKey string) error {
    r._eventKey = _eventKey
    r.Set("event_key", _eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetEventKey() string {
    return r._eventKey
}
// PageNumber Setter
// 第几页
func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageNumber(_pageNumber int64) error {
    r._pageNumber = _pageNumber
    r.Set("page_number", _pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetPageNumber() int64 {
    return r._pageNumber
}
// PageSize Setter
// 分页尺寸
func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetPageSize() int64 {
    return r._pageSize
}
// DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAccessToken(_degAccessToken string) error {
    r._degAccessToken = _degAccessToken
    r.Set("deg_access_token", _degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAccessToken() string {
    return r._degAccessToken
}
