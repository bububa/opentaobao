package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest
天下秀回传订单执行状态变动 API请求
alibaba.pictures.dengta.ims.order.status.change

天下秀回传订单执行状态变动 */
type AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest struct {
	model.Params
	// 状态发生的时间 2020-01-02 10:02:03
	_changeTime string
	// 描述，如ims关单，返回关单原因。
	_comments string
	// 天下秀订单id
	_imsOrderId string
	// 3=抖音，1-微博 2-微信
	_accountType int64
	// 扩展字段
	_extJson string
	// 1:待执行  2:执行中  3:发布  4:完成  5:取消
	_status int64
}

// New
