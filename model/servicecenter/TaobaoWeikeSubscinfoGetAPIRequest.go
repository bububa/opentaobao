package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeSubscinfoGetAPIRequest
需求订单查询接口 API请求
taobao.weike.subscinfo.get

需求订单查询接口 */
type TaobaoWeikeSubscinfoGetAPIRequest struct {
	model.Params
	// 商家旺旺名称
	_sellerName string
	// 时间范围开始时间
	_startTime string
	// 时间范围结束时间
	_endTime string
	// 页码
	_pageNum int64
}

// New
