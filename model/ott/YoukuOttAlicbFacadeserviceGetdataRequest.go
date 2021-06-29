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
    _serviceList   []string
    // 设备唯一标识
    _uuid   string
    // 属性MAP JSON串
    _propertyMapJson   string
    // 扩展属性
    _extraInfoMap   string
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
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetServiceList(_serviceList []string) error {
    r._serviceList = _serviceList
    r.Set("service_list", _serviceList)
    return nil
}

// ServiceList Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetServiceList() []string {
    return r._serviceList
}
// Uuid Setter
// 设备唯一标识
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetUuid() string {
    return r._uuid
}
// PropertyMapJson Setter
// 属性MAP JSON串
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetPropertyMapJson(_propertyMapJson string) error {
    r._propertyMapJson = _propertyMapJson
    r.Set("property_map_json", _propertyMapJson)
    return nil
}

// PropertyMapJson Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetPropertyMapJson() string {
    return r._propertyMapJson
}
// ExtraInfoMap Setter
// 扩展属性
func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetExtraInfoMap(_extraInfoMap string) error {
    r._extraInfoMap = _extraInfoMap
    r.Set("extra_info_map", _extraInfoMap)
    return nil
}

// ExtraInfoMap Getter
func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetExtraInfoMap() string {
    return r._extraInfoMap
}
