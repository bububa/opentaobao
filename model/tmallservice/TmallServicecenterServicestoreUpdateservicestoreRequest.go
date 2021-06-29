package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改网点信息 API请求
tmall.servicecenter.servicestore.updateservicestore

修改网点信息。对于同一个服务商，通过 service_store_code 保证网点唯一性。需要保证网点存在才能修改。
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
type TmallServicecenterServicestoreUpdateservicestoreRequest struct {
    model.Params
    // 网点
    _serviceStore   *ServiceStoreDto
}

// 初始化TmallServicecenterServicestoreUpdateservicestoreRequest对象
func NewTmallServicecenterServicestoreUpdateservicestoreRequest() *TmallServicecenterServicestoreUpdateservicestoreRequest{
    return &TmallServicecenterServicestoreUpdateservicestoreRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreUpdateservicestoreRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.updateservicestore"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreUpdateservicestoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStore Setter
// 网点
func (r *TmallServicecenterServicestoreUpdateservicestoreRequest) SetServiceStore(_serviceStore *ServiceStoreDto) error {
    r._serviceStore = _serviceStore
    r.Set("service_store", _serviceStore)
    return nil
}

// ServiceStore Getter
func (r TmallServicecenterServicestoreUpdateservicestoreRequest) GetServiceStore() *ServiceStoreDto {
    return r._serviceStore
}
