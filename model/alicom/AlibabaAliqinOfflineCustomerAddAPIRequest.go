package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinOfflineCustomerAddAPIRequest
系外拉新代理商增加客户经理接口 API请求
alibaba.aliqin.offline.customer.add

阿里通信这边维护了代理商和其对应的客户经理的关系，用于业务处理，开放该接口用于代理商将他们系统下的客户经理信息同步给我们 */
type AlibabaAliqinOfflineCustomerAddAPIRequest struct {
	model.Params
	// 代理商id
	_distributeId string
	// 网点id，如果存在填写，不存在的话，填0即可；注意：如果填写了这个字段，后面的pob_name等会失效；如果为0，下面的网点名称、省份、城市必填
	_agentId string
	// 网点名称
	_pobName string
	// 网点所在省份
	_province string
	// 网点所在城市
	_city string
	// 客户经理名称
	_customerName string
	// 手机号码
	_phone string
	// 客户经理编码，如果没有可以不填
	_otherKey string
	// json类型，Map<String,String>
	_ext string
	// 活动编码
	_activityCode string
	// token，页面获取到的参数
	_token string
}

// New
