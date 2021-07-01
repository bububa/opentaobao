package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelItemElementManageAPIRequest
【API3.0】资源元素管理接口 API请求
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享 */
type TaobaoAlitripTravelItemElementManageAPIRequest struct {
	model.Params
	// 必填，操作类型：1-新增，2-修改，3-删除。。特别注意：删除 为逻辑删除，即该outer_id所对应的元素还存在但是会置为无效状态，重新编辑修改即可恢复为有效状态。因此该id一旦使用将不可重复
	_operation int64
	// 必填，元素的外部商家编码，必须唯一。编辑、删除时将根据该编码找到对应元素。
	_outerId string
	// 资源元素类型，新增时必填：1-景点，2-酒店，5-交通接驳，6-WIFI库，7-电话卡，8-餐饮，9-签证库，11-特色活动，999-其他
	_elementType int64
	// 元素名称，新增时必填； 注意：Wifi库的使用地和签证库所在国家均适用这个字段
	_name string
	// 元素所在城市，景点、酒店在新增时必填
	_city string
	// 元素的子类型，新增时必填。景点指门票类型，酒店指房型信息，交通接驳（接送机、接驳车、租车、船票、其他）选其一，餐饮（早餐、晚餐、午餐、下午茶及其他）选其一；签证（旅游签证、商务签证、工作签证、留学签证、探亲访友签证、入台证、其他）
	_type string
	// 当新增“交通接驳、餐饮、特色活动、其他”资源类型时 必填
	_desc string
}

// New
