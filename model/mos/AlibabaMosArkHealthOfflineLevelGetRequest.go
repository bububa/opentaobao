package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mall的离线等级 APIRequest
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级
*/
type AlibabaMosArkHealthOfflineLevelGetRequest struct {
    model.Params

    // 商场id
    mallId   string 

}

func NewAlibabaMosArkHealthOfflineLevelGetRequest() *AlibabaMosArkHealthOfflineLevelGetRequest{
    return &AlibabaMosArkHealthOfflineLevelGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetApiMethodName() string {
    return "alibaba.mos.ark.health.offline.level.get"
}

func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosArkHealthOfflineLevelGetRequest) SetMallId(mallId string) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetMallId() string {
    return r.mallId
}

