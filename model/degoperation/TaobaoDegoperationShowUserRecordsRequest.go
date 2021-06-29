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
    degAppKey   string
    // 活动后台配置
    eventKey   string
    // 第几页
    pageNumber   int64
    // 分页尺寸
    pageSize   int64
    // 系统信息
    degAccessToken   string
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
func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

// DegAppKey Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAppKey() string {
    return r.degAppKey
}
// EventKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowUserRecordsRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

// EventKey Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetEventKey() string {
    return r.eventKey
}
// PageNumber Setter
// 第几页
func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetPageNumber() int64 {
    return r.pageNumber
}
// PageSize Setter
// 分页尺寸
func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetPageSize() int64 {
    return r.pageSize
}
// DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAccessToken() string {
    return r.degAccessToken
}
