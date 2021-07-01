package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除网点 API请求
tmall.servicecenter.servicestore.deleteservicestore

删除网点信息。对于同一个服务商，通过 service_store_code 删除网点。
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
type TmallServicecenterServicestoreDeleteservicestoreAPIRequest struct {
    model.Params
    // 网点名称
    _serviceStoreCode   string
}

// 初始化TmallServicecenterServicestoreDeleteservicestoreAPIRequest对象
func NewTmallServicecenterServicestoreDeleteservicestoreRequest() *TmallServicecenterServicestoreDeleteservicestoreAPIRequest{
    return &TmallServicecenterServicestoreDeleteservicestoreAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicestoreDeleteservicestoreAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicestore.deleteservicestore"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicestoreDeleteservicestoreAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceStoreCode Setter
// 网点名称
func (r *TmallServicecenterServicestoreDeleteservicestoreAPIRequest) SetServiceStoreCode(_serviceStoreCode string) error {
    r._serviceStoreCode = _serviceStoreCode
    r.Set("service_store_code", _serviceStoreCode)
    return nil
}

// ServiceStoreCode Getter
func (r TmallServicecenterServicestoreDeleteservicestoreAPIRequest) GetServiceStoreCode() string {
    return r._serviceStoreCode
}
