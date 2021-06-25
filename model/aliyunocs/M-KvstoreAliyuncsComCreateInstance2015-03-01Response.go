package aliyunocs

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建OCS实例 APIResponse
m-kvstore.aliyuncs.com.CreateInstance.2015-03-01

创建OCS实例
*/
type M-KvstoreAliyuncsComCreateInstance2015-03-01APIResponse struct {
    model.CommonResponse
    Response *M-KvstoreAliyuncsComCreateInstance2015-03-01Response `json:"m-kvstore_aliyuncs_com_CreateInstance_2015-03-01_response,omitempty"`
}

type M-KvstoreAliyuncsComCreateInstance2015-03-01Response struct {

    // OCS例ID;<br/>16位字符串(全局唯一)
    InstanceId   string `json:"InstanceId,omitempty"`

    // 实例名称
    InstanceName   string `json:"InstanceName,omitempty"`

    // Memcached连接域名<br/>注：仅支持内网
    ConnectionDomain   string `json:"ConnectionDomain,omitempty"`

    // Memcached连接端口
    Port   int64 `json:"Port,omitempty"`

    // Memcached连接用户名<br/>注：取InstanceId的值
    UserName   string `json:"UserName,omitempty"`

    // 使用中：Available<br/>已禁用：Unavailable<br/>创建中：Creating<br/>修改中：Changing
    InstanceStatus   string `json:"InstanceStatus,omitempty"`

    // cn-hangzhou<br/>cn-qingdao <br/>地域-城市
    RegionId   string `json:"RegionId,omitempty"`

    // 实例容量上限  单位：MByte
    Capacity   int64 `json:"Capacity,omitempty"`

    // 实例QPS限制<br/>单位：次/秒
    QPS   int64 `json:"QPS,omitempty"`

    // 实例带宽限制<br/>单位：MKbps
    Bandwidth   int64 `json:"Bandwidth,omitempty"`

    // 实例连接数限制<br/>单位：个
    Connections   int64 `json:"Connections,omitempty"`

    // 请求唯一标识
    RequestId   string `json:"RequestId,omitempty"`

    // RegionId下级可用区编码
    ZoneId   string `json:"ZoneId,omitempty"`

    // OCS实例的网络类型
    NetworkType   string `json:"NetworkType,omitempty"`

    // NetworkType为vpc时OCS实例所属vpc内由用户指定或系统分配的私网IP地址。NetworkType为classic时忽略此返回参数。
    PrivateIpAddress   string `json:"PrivateIpAddress,omitempty"`

}