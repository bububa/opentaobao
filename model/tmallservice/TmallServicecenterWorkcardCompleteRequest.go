package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单完结 APIRequest
tmall.servicecenter.workcard.complete

工单完结
*/
type TmallServicecenterWorkcardCompleteRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

    // 完结次数
    completeCount   int64 

    // 扩展信息
    extJson   string 

    // 工单完结号
    sequence   int64 

    // 核销地纬度
    latitude   string 

    // 核销地经度
    longitude   string 

}

func NewTmallServicecenterWorkcardCompleteRequest() *TmallServicecenterWorkcardCompleteRequest{
    return &TmallServicecenterWorkcardCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardCompleteRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.complete"
}

func (r TmallServicecenterWorkcardCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardCompleteRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardCompleteRequest) SetCompleteCount(completeCount int64) error {
    r.completeCount = completeCount
    r.Set("complete_count", completeCount)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetCompleteCount() int64 {
    return r.completeCount
}

func (r *TmallServicecenterWorkcardCompleteRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetExtJson() string {
    return r.extJson
}

func (r *TmallServicecenterWorkcardCompleteRequest) SetSequence(sequence int64) error {
    r.sequence = sequence
    r.Set("sequence", sequence)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetSequence() int64 {
    return r.sequence
}

func (r *TmallServicecenterWorkcardCompleteRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetLatitude() string {
    return r.latitude
}

func (r *TmallServicecenterWorkcardCompleteRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r TmallServicecenterWorkcardCompleteRequest) GetLongitude() string {
    return r.longitude
}

