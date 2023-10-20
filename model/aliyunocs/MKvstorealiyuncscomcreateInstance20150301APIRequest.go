package aliyunocs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// MKvstoreAliyuncsComCreateInstance20150301APIRequest 创建OCS实例 API请求
// m-kvstore.aliyuncs.com.CreateInstance.2015-03-01
//
// 创建OCS实例
type MKvstoreAliyuncsComCreateInstance20150301APIRequest struct {
	model.Params
	// 实例昵称长度为2-128个字符，以大小写英文字母或中文开头，不支持字符@/:="<>{[]}和空格
	_instanceName string
	// 说明：实例密码规则：需同时且只能包含大写字母、小写字母和数字长度8~30个字符
	_password string
	// 华东杭州：cn-hangzhou 华北青岛：cn-qingdao
	_regionId string
	// 用于保证幂等性
	_token string
	// RegionId下级可用区编码
	_zoneId string
	// OCS实例的网络类型：classic或vpc。默认为classic
	_networkType string
	// OCS实例所属vpc实例id。如果NetworkType参数为vpc，则本参数为必须。
	_vpcId string
	// OCS实例所属vpc实例的虚拟交换机ID。如果NetworkType参数为vpc，则本参数为必须。
	_vSwitchId string
	// OCS实例所属vpc内的私网IP地址。当NetworkType参数为vpc时有效。如果不带本参数，则系统通过VpcId和VSwitchId自动分配。如果此参数对应的私网IP地址不在VSwitchId包含的IP地址段内，则创建失败。
	_privateIpAddress string
	// 说明：实例容量设置单位：MByte输入范围请见OCS规格参数表
	_capacity int64
}

// NewMKvstoreAliyuncsComCreateInstance20150301Request 初始化MKvstoreAliyuncsComCreateInstance20150301APIRequest对象
func NewMKvstoreAliyuncsComCreateInstance20150301Request() *MKvstoreAliyuncsComCreateInstance20150301APIRequest {
	return &MKvstoreAliyuncsComCreateInstance20150301APIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) Reset() {
	r._instanceName = ""
	r._password = ""
	r._regionId = ""
	r._token = ""
	r._zoneId = ""
	r._networkType = ""
	r._vpcId = ""
	r._vSwitchId = ""
	r._privateIpAddress = ""
	r._capacity = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetApiMethodName() string {
	return "m-kvstore.aliyuncs.com.CreateInstance.2015-03-01"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInstanceName is InstanceName Setter
// 实例昵称长度为2-128个字符，以大小写英文字母或中文开头，不支持字符@/:=&#34;&lt;&gt;{[]}和空格
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetInstanceName(_instanceName string) error {
	r._instanceName = _instanceName
	r.Set("InstanceName", _instanceName)
	return nil
}

// GetInstanceName InstanceName Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetInstanceName() string {
	return r._instanceName
}

// SetPassword is Password Setter
// 说明：实例密码规则：需同时且只能包含大写字母、小写字母和数字长度8~30个字符
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetPassword(_password string) error {
	r._password = _password
	r.Set("Password", _password)
	return nil
}

// GetPassword Password Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetPassword() string {
	return r._password
}

// SetRegionId is RegionId Setter
// 华东杭州：cn-hangzhou 华北青岛：cn-qingdao
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetRegionId(_regionId string) error {
	r._regionId = _regionId
	r.Set("RegionId", _regionId)
	return nil
}

// GetRegionId RegionId Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetRegionId() string {
	return r._regionId
}

// SetToken is Token Setter
// 用于保证幂等性
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("Token", _token)
	return nil
}

// GetToken Token Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetToken() string {
	return r._token
}

// SetZoneId is ZoneId Setter
// RegionId下级可用区编码
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetZoneId(_zoneId string) error {
	r._zoneId = _zoneId
	r.Set("ZoneId", _zoneId)
	return nil
}

// GetZoneId ZoneId Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetZoneId() string {
	return r._zoneId
}

// SetNetworkType is NetworkType Setter
// OCS实例的网络类型：classic或vpc。默认为classic
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetNetworkType(_networkType string) error {
	r._networkType = _networkType
	r.Set("NetworkType", _networkType)
	return nil
}

// GetNetworkType NetworkType Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetNetworkType() string {
	return r._networkType
}

// SetVpcId is VpcId Setter
// OCS实例所属vpc实例id。如果NetworkType参数为vpc，则本参数为必须。
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetVpcId(_vpcId string) error {
	r._vpcId = _vpcId
	r.Set("VpcId", _vpcId)
	return nil
}

// GetVpcId VpcId Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetVpcId() string {
	return r._vpcId
}

// SetVSwitchId is VSwitchId Setter
// OCS实例所属vpc实例的虚拟交换机ID。如果NetworkType参数为vpc，则本参数为必须。
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetVSwitchId(_vSwitchId string) error {
	r._vSwitchId = _vSwitchId
	r.Set("VSwitchId", _vSwitchId)
	return nil
}

// GetVSwitchId VSwitchId Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetVSwitchId() string {
	return r._vSwitchId
}

// SetPrivateIpAddress is PrivateIpAddress Setter
// OCS实例所属vpc内的私网IP地址。当NetworkType参数为vpc时有效。如果不带本参数，则系统通过VpcId和VSwitchId自动分配。如果此参数对应的私网IP地址不在VSwitchId包含的IP地址段内，则创建失败。
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetPrivateIpAddress(_privateIpAddress string) error {
	r._privateIpAddress = _privateIpAddress
	r.Set("PrivateIpAddress", _privateIpAddress)
	return nil
}

// GetPrivateIpAddress PrivateIpAddress Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetPrivateIpAddress() string {
	return r._privateIpAddress
}

// SetCapacity is Capacity Setter
// 说明：实例容量设置单位：MByte输入范围请见OCS规格参数表
func (r *MKvstoreAliyuncsComCreateInstance20150301APIRequest) SetCapacity(_capacity int64) error {
	r._capacity = _capacity
	r.Set("Capacity", _capacity)
	return nil
}

// GetCapacity Capacity Getter
func (r MKvstoreAliyuncsComCreateInstance20150301APIRequest) GetCapacity() int64 {
	return r._capacity
}

var poolMKvstoreAliyuncsComCreateInstance20150301APIRequest = sync.Pool{
	New: func() any {
		return NewMKvstoreAliyuncsComCreateInstance20150301Request()
	},
}

// GetMKvstoreAliyuncsComCreateInstance20150301Request 从 sync.Pool 获取 MKvstoreAliyuncsComCreateInstance20150301APIRequest
func GetMKvstoreAliyuncsComCreateInstance20150301APIRequest() *MKvstoreAliyuncsComCreateInstance20150301APIRequest {
	return poolMKvstoreAliyuncsComCreateInstance20150301APIRequest.Get().(*MKvstoreAliyuncsComCreateInstance20150301APIRequest)
}

// ReleaseMKvstoreAliyuncsComCreateInstance20150301APIRequest 将 MKvstoreAliyuncsComCreateInstance20150301APIRequest 放入 sync.Pool
func ReleaseMKvstoreAliyuncsComCreateInstance20150301APIRequest(v *MKvstoreAliyuncsComCreateInstance20150301APIRequest) {
	v.Reset()
	poolMKvstoreAliyuncsComCreateInstance20150301APIRequest.Put(v)
}
