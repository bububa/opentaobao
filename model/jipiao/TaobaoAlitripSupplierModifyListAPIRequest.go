package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSupplierModifyListAPIRequest
【机票代理商订单】改签通知单列表 API请求
taobao.alitrip.supplier.modify.list

提供供应商查询改签通知单列表 */
type TaobaoAlitripSupplierModifyListAPIRequest struct {
	model.Params
	// 页码
	_currentPage int64
	// 乘客出发时间查询结束日期
	_depEnd string
	// 乘客出发时间查询开始日期
	_depStart string
	// 申请单创建时间查询结束日期
	_gmtCreateEnd string
	// 申请单创建时间查询开始日期
	_gmtCreateStart string
	// 每页记录数
	_pageSize int64
	// 1：改签申请列表，2：改签已支付列表，3：改签成功列表
	_status int64
}

// New
