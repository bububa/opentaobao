package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketTasksGetAPIRequest
任务列表获取接口 API请求
taobao.vmarket.eticket.tasks.get

外部合作卖家获取任务列表的信息：如发码同通知失败或者回调失败的订单号 */
type TaobaoVmarketEticketTasksGetAPIRequest struct {
	model.Params
	// 卖家家ID(信任卖家不必传，码商可选)
	_sellerId int64
	// 返回结果类型:<br/>1:返回通知失败的订单<br/>2.返回通知成功回调失败的订单
	_type int64
	// 页码。取值范围:大于零的整数; 默认值:1
	_pageNo int64
	// 每页获取条数。默认值40，最小值1，最大值100。
	_pageSize int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
}

// New
