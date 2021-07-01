package aliyunocs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* M_kvstoreAliyuncsComCreateInstance2015_03_01APIRequest
创建OCS实例 API请求
m-kvstore.aliyuncs.com.CreateInstance.2015-03-01

创建OCS实例 */
type M_kvstoreAliyuncsComCreateInstance2015_03_01APIRequest struct {
	model.Params
	// 实例昵称长度为2-128个字符，以大小写英文字母或中文开头，不支持字符@/:="<>{[]}和空格
	_instanceName string
	// 说明：实例密码规则：需同时且只能包含大写字母、小写字母和数字长度8~30个字符
	_password string
	// 说明：实例容量设置单位：MByte输入范围请见OCS规格参数表
	_capacity int64
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
}

// New
