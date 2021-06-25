package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务网点创建 APIRequest
tmall.servicecenter.servicestore.createservicestore

创建网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点不存在才能创建。地址信息：中文和编码二选一，都填则以编码address_code为准。
错误码
1, 服务商昵称无效
2, 缺少省份
3, 缺少城市
4, 缺少区域
5, 缺少详细地址
6, 传入地址不在标准地址库中，无法解析
7, 缺少网点编码
8, 缺少网点名称
9, 缺少网点电话
10, 网点已存在
11, 网点不存在
12, 系统错误
*/
type TmallServicecenterServicestoreCreateservicestoreRequest struct {
    model.Params

    // 网点
    serviceStore   *ServiceStoreDTO 

}

func NewTmallServicecenterServicestoreCreateservicestoreRequest() *TmallServicecenterServicestoreCreateservicestoreRequest{
    return &TmallServicecenterServicestoreCreateservicestoreRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicestoreCreateservicestoreRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.createservicestore"
}

func (r TmallServicecenterServicestoreCreateservicestoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicestoreCreateservicestoreRequest) SetServiceStore(serviceStore *ServiceStoreDTO) error {
    r.serviceStore = serviceStore
    r.Set("service_store", serviceStore)
    return nil
}

func (r TmallServicecenterServicestoreCreateservicestoreRequest) GetServiceStore() *ServiceStoreDTO {
    return r.serviceStore
}

