package aliyunocs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OCS实例 APIRequest
m-kvstore.aliyuncs.com.CreateInstance.2015-03-01

创建OCS实例
*/
type M-KvstoreAliyuncsComCreateInstance2015-03-01Request struct {
    model.Params

    // 实例昵称长度为2-128个字符，以大小写英文字母或中文开头，不支持字符@/:="<>{[]}和空格
    instanceName   string 

    // 说明：实例密码规则：需同时且只能包含大写字母、小写字母和数字长度8~30个字符
    password   string 

    // 说明：实例容量设置单位：MByte输入范围请见OCS规格参数表
    capacity   int64 

    // 华东杭州：cn-hangzhou 华北青岛：cn-qingdao
    regionId   string 

    // 用于保证幂等性
    token   string 

    // RegionId下级可用区编码
    zoneId   string 

    // OCS实例的网络类型：classic或vpc。默认为classic
    networkType   string 

    // OCS实例所属vpc实例id。如果NetworkType参数为vpc，则本参数为必须。
    vpcId   string 

    // OCS实例所属vpc实例的虚拟交换机ID。如果NetworkType参数为vpc，则本参数为必须。
    vSwitchId   string 

    // OCS实例所属vpc内的私网IP地址。当NetworkType参数为vpc时有效。如果不带本参数，则系统通过VpcId和VSwitchId自动分配。如果此参数对应的私网IP地址不在VSwitchId包含的IP地址段内，则创建失败。
    privateIpAddress   string 

}

func NewM-KvstoreAliyuncsComCreateInstance2015-03-01Request() *M-KvstoreAliyuncsComCreateInstance2015-03-01Request{
    return &M-KvstoreAliyuncsComCreateInstance2015-03-01Request{
        Params: model.NewParams(),
    }
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetApiMethodName() string {
    return "m-kvstore.aliyuncs.com.CreateInstance.2015-03-01"
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("InstanceName", instanceName)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetInstanceName() string {
    return r.instanceName
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetPassword(password string) error {
    r.password = password
    r.Set("Password", password)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetPassword() string {
    return r.password
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetCapacity(capacity int64) error {
    r.capacity = capacity
    r.Set("Capacity", capacity)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetCapacity() int64 {
    return r.capacity
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetRegionId(regionId string) error {
    r.regionId = regionId
    r.Set("RegionId", regionId)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetRegionId() string {
    return r.regionId
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetToken(token string) error {
    r.token = token
    r.Set("Token", token)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetToken() string {
    return r.token
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetZoneId(zoneId string) error {
    r.zoneId = zoneId
    r.Set("ZoneId", zoneId)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetZoneId() string {
    return r.zoneId
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetNetworkType(networkType string) error {
    r.networkType = networkType
    r.Set("NetworkType", networkType)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetNetworkType() string {
    return r.networkType
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetVpcId(vpcId string) error {
    r.vpcId = vpcId
    r.Set("VpcId", vpcId)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetVpcId() string {
    return r.vpcId
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetVSwitchId(vSwitchId string) error {
    r.vSwitchId = vSwitchId
    r.Set("VSwitchId", vSwitchId)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetVSwitchId() string {
    return r.vSwitchId
}

func (r *M-KvstoreAliyuncsComCreateInstance2015-03-01Request) SetPrivateIpAddress(privateIpAddress string) error {
    r.privateIpAddress = privateIpAddress
    r.Set("PrivateIpAddress", privateIpAddress)
    return nil
}

func (r M-KvstoreAliyuncsComCreateInstance2015-03-01Request) GetPrivateIpAddress() string {
    return r.privateIpAddress
}

