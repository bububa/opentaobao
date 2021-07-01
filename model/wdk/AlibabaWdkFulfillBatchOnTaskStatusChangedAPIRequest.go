package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest
物流管控作业状态回传 API请求
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传 */
type AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest struct {
	model.Params
	// 作业状态回传对象
	_taskStatus *TaskStatus
}

// New
