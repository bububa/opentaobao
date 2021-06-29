package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单完结 API请求
tmall.servicecenter.workcard.complete

工单完结
*/
type TmallServicecenterWorkcardCompleteRequest struct {
    model.Params
    // 工单id
    _workcardId   int64
    // 完结次数
    _completeCount   int64
    // 扩展信息
    _extJson   string
    // 工单完结号
    _sequence   int64
    // 核销地纬度
    _latitude   string
    // 核销地经度
    _longitude   string
}

// 初始化TmallServicecenterWorkcardCompleteRequest对象
func NewTmallServicecenterWorkcardCompleteRequest() *TmallServicecenterWorkcardCompleteRequest{
    return &TmallServicecenterWorkcardCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCompleteRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.complete"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardCompleteRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// CompleteCount Setter
// 完结次数
func (r *TmallServicecenterWorkcardCompleteRequest) SetCompleteCount(_completeCount int64) error {
    r._completeCount = _completeCount
    r.Set("complete_count", _completeCount)
    return nil
}

// CompleteCount Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetCompleteCount() int64 {
    return r._completeCount
}
// ExtJson Setter
// 扩展信息
func (r *TmallServicecenterWorkcardCompleteRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetExtJson() string {
    return r._extJson
}
// Sequence Setter
// 工单完结号
func (r *TmallServicecenterWorkcardCompleteRequest) SetSequence(_sequence int64) error {
    r._sequence = _sequence
    r.Set("sequence", _sequence)
    return nil
}

// Sequence Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetSequence() int64 {
    return r._sequence
}
// Latitude Setter
// 核销地纬度
func (r *TmallServicecenterWorkcardCompleteRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetLatitude() string {
    return r._latitude
}
// Longitude Setter
// 核销地经度
func (r *TmallServicecenterWorkcardCompleteRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TmallServicecenterWorkcardCompleteRequest) GetLongitude() string {
    return r._longitude
}
