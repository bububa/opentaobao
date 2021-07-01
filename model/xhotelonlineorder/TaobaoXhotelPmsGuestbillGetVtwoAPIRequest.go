package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelPmsGuestbillGetVtwoAPIRequest
客人PMS账单信息查询 API请求
taobao.xhotel.pms.guestbill.get.vtwo

从pms获取客人账单信息 */
type TaobaoXhotelPmsGuestbillGetVtwoAPIRequest struct {
	model.Params
	// 开票点税号
	_taxNum string
	// 身份证后4位
	_shortIdNum string
	// 房间号
	_roomNum string
	// 请求id (32位唯一值)
	_requestId string
	// 用户渠道(0:未知,1:淘宝)
	_userChannel int64
}

// New
