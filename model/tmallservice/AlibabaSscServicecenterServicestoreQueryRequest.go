package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店信息 API请求
alibaba.ssc.servicecenter.servicestore.query

根据天猫id查询门店信息
*/
type AlibabaSscServicecenterServicestoreQueryRequest struct {
    model.Params
    // 天猫id
    _id   int64
}

// 初始化AlibabaSscServicecenterServicestoreQueryRequest对象
func NewAlibabaSscServicecenterServicestoreQueryRequest() *AlibabaSscServicecenterServicestoreQueryRequest{
    return &AlibabaSscServicecenterServicestoreQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscServicecenterServicestoreQueryRequest) GetApiMethodName() string {
    return "alibaba.ssc.servicecenter.servicestore.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscServicecenterServicestoreQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 天猫id
func (r *AlibabaSscServicecenterServicestoreQueryRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaSscServicecenterServicestoreQueryRequest) GetId() int64 {
    return r._id
}
