package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasOrderSearchAPIRequest
订单记录导出 API请求
taobao.vas.order.search

用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。<br/>建议用于查询前一日的历史记录，不适合用作实时数据查询。<br/>现在只能查询90天以内的数据<br/>该接口限制每分钟所有appkey调用总和只能有800次。 */
type TaobaoVasOrderSearchAPIRequest struct {
	model.Params
	// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
	_articleCode string
	// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
	_itemCode string
	// 淘宝会员名
	_nick string
	// 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
	_startCreated string
	// 订单创建时间（订购时间）结束值
	_endCreated string
	// 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
	_bizType int64
	// 订单号
	_bizOrderId int64
	// 子订单号
	_orderId int64
	// 一页包含的记录数
	_pageSize int64
	// 页码
	_pageNo int64
}

// New
