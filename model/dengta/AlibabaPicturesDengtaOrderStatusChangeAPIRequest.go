package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaOrderStatusChangeAPIRequest
天下秀订单状态变更通知 API请求
alibaba.pictures.dengta.order.status.change

天下秀订单状态变更通知 */
type AlibabaPicturesDengtaOrderStatusChangeAPIRequest struct {
	model.Params
	// 拒绝原因
	_remark string
	// 新状态
	_status int64
	// 变更时间
	_changeTime string
	// ims订单编号
	_imsOrderId string
	// task 编号
	_aliTaskId string
}

// New
