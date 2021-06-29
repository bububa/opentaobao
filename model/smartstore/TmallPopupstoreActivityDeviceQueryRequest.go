package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据活动id查询活动相关快闪店及设备信息 APIRequest
tmall.popupstore.activity.device.query

查询某一活动的deviceCode的部署情况
*/
type TmallPopupstoreActivityDeviceQueryRequest struct {
    model.Params

    // ISV的活动ID
    activityId   int64 

}

func NewTmallPopupstoreActivityDeviceQueryRequest() *TmallPopupstoreActivityDeviceQueryRequest{
    return &TmallPopupstoreActivityDeviceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TmallPopupstoreActivityDeviceQueryRequest) GetApiMethodName() string {
    return "tmall.popupstore.activity.device.query"
}

func (r TmallPopupstoreActivityDeviceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallPopupstoreActivityDeviceQueryRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TmallPopupstoreActivityDeviceQueryRequest) GetActivityId() int64 {
    return r.activityId
}

