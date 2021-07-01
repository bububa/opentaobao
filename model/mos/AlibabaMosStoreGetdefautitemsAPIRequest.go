package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosStoreGetdefautitemsAPIRequest
获取默认状态下商品列表 API请求
alibaba.mos.store.getdefautitems

获取默认状态下商品列表 */
type AlibabaMosStoreGetdefautitemsAPIRequest struct {
	model.Params
	// 屏编号
	_screenNo string
	// 分页查询开始index
	_start int64
	// 分页查询每页记录数
	_limitCount int64
}

// New
