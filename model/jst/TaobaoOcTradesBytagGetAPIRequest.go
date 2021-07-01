package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcTradesBytagGetAPIRequest
标签查询订单 API请求
taobao.oc.trades.bytag.get

根据标签查询订单编号 */
type TaobaoOcTradesBytagGetAPIRequest struct {
	model.Params
	// 分页大小
	_pageSize int64
	// 标签类型，1官方，2自定义
	_tagType int64
	// 当前页
	_page int64
	// 标签名称
	_tagName string
}

// New
