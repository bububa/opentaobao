package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户中奖记录 APIRequest
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

func NewTaobaoDegoperationShowUserRecordsRequest() *TaobaoDegoperationShowUserRecordsRequest{
    return &TaobaoDegoperationShowUserRecordsRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetApiMethodName() string {
    return "taobao.degoperation.show.user.records"
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAppKey(degAppKey string) error {
    r.degAppKey = degAppKey
    r.Set("deg_app_key", degAppKey)
    return nil
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAppKey() string {
    return r.degAppKey
}

func (r *TaobaoDegoperationShowUserRecordsRequest) SetEventKey(eventKey string) error {
    r.eventKey = eventKey
    r.Set("event_key", eventKey)
    return nil
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetEventKey() string {
    return r.eventKey
}

func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetPageNumber() int64 {
    return r.pageNumber
}

func (r *TaobaoDegoperationShowUserRecordsRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoDegoperationShowUserRecordsRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

func (r TaobaoDegoperationShowUserRecordsRequest) GetDegAccessToken() string {
    return r.degAccessToken
}

