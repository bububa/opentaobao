package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerModifyListAPIRequest
【机票代理商订单】改签订单列表 API请求
taobao.alitrip.seller.modify.list

提供机票代理商查询改签订单列表 */
type TaobaoAlitripSellerModifyListAPIRequest struct {
	model.Params
	// 申请单ID
	_applyId int64
	// 淘宝订单号
	_orderId int64
	// 改签发起时间的查询结束日期 和 更新时间必选其一
	_applyDateEnd string
	// 改签发起时间的查询开始日期 和 更新时间必选其一
	_applyDateStart string
	// 页码
	_currentPage int64
	// 乘客起飞时间的查询结束日期
	_flyDateEnd string
	// 乘客起飞时间的查询开始日期
	_flyDateStart string
	// 每页记录数
	_pageSize int64
	// 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
	_status int64
	// 记录修改结束时间  和 改签发起时间必选其一
	_modifyDateEnd string
	// 记录修改起始时间 和 改签发起时间必选其一
	_modifyDateStart string
}

// New
