package westcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商场列表 APIRequest
alibaba.westcrm.mall.list.get

根据园区id获取商场列表
*/
type AlibabaWestcrmMallListGetRequest struct {
    model.Params

    // 园区id
    campusId   int64 

}

func NewAlibabaWestcrmMallListGetRequest() *AlibabaWestcrmMallListGetRequest{
    return &AlibabaWestcrmMallListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWestcrmMallListGetRequest) GetApiMethodName() string {
    return "alibaba.westcrm.mall.list.get"
}

func (r AlibabaWestcrmMallListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWestcrmMallListGetRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaWestcrmMallListGetRequest) GetCampusId() int64 {
    return r.campusId
}

