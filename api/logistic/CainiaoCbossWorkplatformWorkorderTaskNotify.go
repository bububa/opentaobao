package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* CainiaoCbossWorkplatformWorkorderTaskNotify
TOP-SPI工单任务下发接口
cainiao.cboss.workplatform.workorder.task.notify

TOP-SPI工单任务下发接口（菜鸟--->商家ISV） */
func CainiaoCbossWorkplatformWorkorderTaskNotify(clt *core.SDKClient, req *logistic.CainiaoCbossWorkplatformWorkorderTaskNotifyAPIRequest, session string) (*logistic.CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse, error) {
	var resp logistic.CainiaoCbossWorkplatformWorkorderTaskNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
