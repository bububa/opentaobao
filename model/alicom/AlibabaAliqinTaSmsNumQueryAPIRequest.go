package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTaSmsNumQueryAPIRequest
短信查询 API请求
alibaba.aliqin.ta.sms.num.query

查询短信发送揭露 */
type AlibabaAliqinTaSmsNumQueryAPIRequest struct {
	model.Params
	// 短信发送流水
	_bizId string
	// 短信接收号码
	_recNum string
	// 短信发送日期，支持近30天记录查询，格式yyyyMMdd
	_queryDate string
	// 分页参数,页码
	_currentPage int64
	// 分页参数，每页数量。最大值50
	_pageSize int64
}

// New
