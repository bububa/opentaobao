package cmns

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosServiceCmnsCoaMessageAcksListAPIRequest
消息ack记录查询 API请求
yunos.service.cmns.coa.message.acks.list

第三方应用开发者调用此接口查询消息ack记录 */
type YunosServiceCmnsCoaMessageAcksListAPIRequest struct {
	model.Params
	// 消息id
	_mid int64
	// 设备id
	_did int64
	// 分页查询页码
	_pageIndex int64
	// 分页每页数据集数
	_pageSize int64
}

// New
