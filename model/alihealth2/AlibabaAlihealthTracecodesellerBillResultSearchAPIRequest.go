package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest
查询出入库单处理结果api API请求
alibaba.alihealth.tracecodeseller.bill.result.search

查询出入库单处理结果api */
type AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest struct {
	model.Params
	// top身份认证
	_skeyCode string
	// 商家id
	_entInfoId int64
	// 单据编号
	_billCode string
	// 查询开始日期
	_beginDate string
	// 查询结束日期
	_endDate string
	// 不需要
	_sellerName string
	// 每页条数
	_pageSize int64
	// 当前页
	_page int64
}

// New
