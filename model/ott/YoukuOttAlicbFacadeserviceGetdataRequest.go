package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
影视SDK获取设备能力值 APIRequest
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

func NewYoukuOttAlicbFacadeserviceGetdataRequest() *YoukuOttAlicbFacadeserviceGetdataRequest{
    return &YoukuOttAlicbFacadeserviceGetdataRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetApiMethodName() string {
    return "youku.ott.alicb.facadeservice.getdata"
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetServiceList(serviceList []string) error {
    r.serviceList = serviceList
    r.Set("service_list", serviceList)
    return nil
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetServiceList() []string {
    return r.serviceList
}

func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetUuid() string {
    return r.uuid
}

func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetPropertyMapJson(propertyMapJson string) error {
    r.propertyMapJson = propertyMapJson
    r.Set("property_map_json", propertyMapJson)
    return nil
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetPropertyMapJson() string {
    return r.propertyMapJson
}

func (r *YoukuOttAlicbFacadeserviceGetdataRequest) SetExtraInfoMap(extraInfoMap string) error {
    r.extraInfoMap = extraInfoMap
    r.Set("extra_info_map", extraInfoMap)
    return nil
}

func (r YoukuOttAlicbFacadeserviceGetdataRequest) GetExtraInfoMap() string {
    return r.extraInfoMap
}

