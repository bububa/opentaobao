package eleenterpriseemployee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

// Alibabaeleenterpriseemployeebatchupdate 批量新增更新员工
// alibaba.ele.enterprise.employee.batchupdate
//
// 批量新增更新员工
func Alibabaeleenterpriseemployeebatchupdate(clt *core.SDKClient, req *eleenterpriseemployee.AlibabaeleenterpriseemployeebatchupdateAPIRequest, session string) (*eleenterpriseemployee.AlibabaeleenterpriseemployeebatchupdateAPIResponse, error) {
	var resp eleenterpriseemployee.AlibabaeleenterpriseemployeebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
