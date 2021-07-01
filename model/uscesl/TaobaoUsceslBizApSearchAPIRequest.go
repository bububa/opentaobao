package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizApSearchAPIRequest
AP列表查询 API请求
taobao.uscesl.biz.ap.search

查询当前门店下登记的AP列表 */
type TaobaoUsceslBizApSearchAPIRequest struct {
	model.Params
	// 商家编码
	_bizBrandKey string
	// 每页显示数
	_limit int64
	// 是否激活
	_isActivate bool
	// 价签条码
	_mac string
	// 页码
	_currentPage int64
	// 门店ID
	_storeId int64
}

// New
