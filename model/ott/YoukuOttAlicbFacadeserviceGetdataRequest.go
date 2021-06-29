package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
影视SDK获取设备能力值 API请求
youku.ott.alicb.facadeservice.getdata

影视SDK获取设备能力值
*/
type YoukuOttAlicbFacadeserviceGetdataRequest struct {
    model.Params
    // 能力维度
    serviceList   []string
    // 设备唯一标识
    uuid   string
    // 属性MAP JSON串
    propertyMapJson   string
    // 扩展属性
    extraInfoMap   string
}

// 初始化YoukuOttAlicbFacadeserviceGetdataRequest对象
func NewYoukuOttAlicbFacadeserviceGetdataRequest() *YoukuOttAlicbFacadeserviceGetdataRequest{
    return &YoukuOttAlicbFacadeserviceGetdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetApiMethodName() string {
    return "youku.ott.alicb.facadeservice.getdata"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceList Setter
// 能力维度
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetServiceList(serviceList []string) error {
    r.serviceList = serviceList
    r.Set("service_list", serviceList)
    return nil
}

// ServiceList Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetServiceList() []string {
    return r.serviceList
}
// Uuid Setter
// 设备唯一标识
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetUuid() string {
    return r.uuid
}
// PropertyMapJson Setter
// 属性MAP JSON串
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetPropertyMapJson(propertyMapJson string) error {
    r.propertyMapJson = propertyMapJson
    r.Set("property_map_json", propertyMapJson)
    return nil
}

// PropertyMapJson Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetPropertyMapJson() string {
    return r.propertyMapJson
}
// ExtraInfoMap Setter
// 扩展属性
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetExtraInfoMap(extraInfoMap string) error {
    r.extraInfoMap = extraInfoMap
    r.Set("extra_info_map", extraInfoMap)
    return nil
}

// ExtraInfoMap Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetExtraInfoMap() string {
    return r.extraInfoMap
}
