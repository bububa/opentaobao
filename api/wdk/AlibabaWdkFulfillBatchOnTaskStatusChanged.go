package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkFulfillBatchOnTaskStatusChanged
物流管控作业状态回传
alibaba.wdk.fulfill.batch.on.task.status.changed

物流管控作业状态回传 */
func AlibabaWdkFulfillBatchOnTaskStatusChanged(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIRequest, session string) (*wdk.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse, error) {
	var resp wdk.AlibabaWdkFulfillBatchOnTaskStatusChangedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
