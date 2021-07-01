package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallExchangeAgreeAPIRequest
卖家同意换货申请 API请求
tmall.exchange.agree

卖家同意换货申请 */
type TmallExchangeAgreeAPIRequest struct {
	model.Params
	// 邮政编码
	_post string
	// 上传图片举证
	_leaveMessagePics *model.File
	// 卖家留言
	_leaveMessage string
	// 收货地址id，如需获取请调用该top接口：taobao.logistics.address.search，对应属性为contact_id
	_addressId int64
	// 详细收货地址
	_completeAddress string
	// 换货单号ID
	_disputeId int64
	// 返回字段。当前支持的有 dispute_id, bizorder_id, modified, status
	_fields []string
	// 收货人手机号
	_mobile string
}

// New
