package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsPartnersGetAPIRequest
查询支持起始地到目的地范围的物流公司 API请求
taobao.logistics.partners.get

查询物流公司信息（可以查询目的地可不可达情况） */
type TaobaoLogisticsPartnersGetAPIRequest struct {
	model.Params
	// 物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
	_sourceId string
	// 物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
	_targetId string
	// 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。
	_serviceType string
	// 货物价格.只有当选择货到付款此参数才会有效
	_goodsValue string
	// 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。
	_isNeedCarriage bool
}

// New
