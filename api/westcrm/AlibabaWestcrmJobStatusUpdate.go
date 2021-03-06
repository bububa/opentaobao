package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmJobStatusUpdate 更新工单状态
// alibaba.westcrm.job.status.update
//
// 更新工单处理状态
func AlibabaWestcrmJobStatusUpdate(clt *core.SDKClient, req *westcrm.AlibabaWestcrmJobStatusUpdateAPIRequest, session string) (*westcrm.AlibabaWestcrmJobStatusUpdateAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmJobStatusUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
