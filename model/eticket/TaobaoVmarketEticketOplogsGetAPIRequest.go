package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketOplogsGetAPIRequest
电子凭证操作日志查询 API请求
taobao.vmarket.eticket.oplogs.get

电子凭证核销日志查询 */
type TaobaoVmarketEticketOplogsGetAPIRequest struct {
	model.Params
	// 0:全部 1:核销 2:冲正
	_type int64
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 核销码
	_code string
	// 手机号后四位
	_mobile string
	// 当前页码
	_pageNo int64
	// 每页显示的记录数，最大为40，默认为40
	_pageSize int64
	// 排序方式
	_sort string
	// 核销身份
	_posid string
	// 码商ID
	_codemerchantId int64
}

// New
