package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCbossWorkplatformOperationReplyAPIRequest
菜鸟工单操作回传 API请求
cainiao.cboss.workplatform.operation.reply

菜鸟工单进度下发接口，目前调用者ISV */
type CainiaoCbossWorkplatformOperationReplyAPIRequest struct {
	model.Params
	// 工单id
	_workOrderId string
	// 工单任务id
	_taskId string
	// 任务操作时间
	_actionTime string
	// 任务操作类型
	_actionType int64
	// 操作者userId
	_dealerUserId string
	// 操作者联系方式
	_dealerContact string
	// 商家工单操作回传备注
	_memo string
	// 凭证照片地址拼接
	_attachPath string
	// 扩展字段
	_features string
}

// New
