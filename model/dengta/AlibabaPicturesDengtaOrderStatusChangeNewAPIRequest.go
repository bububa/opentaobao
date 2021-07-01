package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest
天下秀订单状态变更通知 API请求
alibaba.pictures.dengta.order.status.change.new

天下秀订单状态变更通知 */
type AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest struct {
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
	// 发布内容
	_taskContent string
	// 发布图片url列表
	_taskPic string
	// 扩展字段。json结构
	_extJson string
}

// New
