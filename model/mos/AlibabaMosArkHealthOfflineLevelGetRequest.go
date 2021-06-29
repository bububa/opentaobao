package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mall的离线等级 API请求
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级
*/
type AlibabaMosArkHealthOfflineLevelGetRequest struct {
    model.Params
    // 商场id
    _mallId   string
}

// 初始化AlibabaMosArkHealthOfflineLevelGetRequest对象
func NewAlibabaMosArkHealthOfflineLevelGetRequest() *AlibabaMosArkHealthOfflineLevelGetRequest{
    return &AlibabaMosArkHealthOfflineLevelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetApiMethodName() string {
    return "alibaba.mos.ark.health.offline.level.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MallId Setter
// 商场id
func (r *AlibabaMosArkHealthOfflineLevelGetRequest) SetMallId(_mallId string) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r AlibabaMosArkHealthOfflineLevelGetRequest) GetMallId() string {
    return r._mallId
}
