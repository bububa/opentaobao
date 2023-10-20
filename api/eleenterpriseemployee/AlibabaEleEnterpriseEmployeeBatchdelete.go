package eleenterpriseemployee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseemployee"
)

// Alibabaeleenterpriseemployeebatchdelete 批量删除员工
// alibaba.ele.enterprise.employee.batchdelete
//
// 批量删除员工
func Alibabaeleenterpriseemployeebatchdelete(clt *core.SDKClient, req *eleenterpriseemployee.AlibabaeleenterpriseemployeebatchdeleteAPIRequest, session string) (*eleenterpriseemployee.AlibabaeleenterpriseemployeebatchdeleteAPIResponse, error) {
	var resp eleenterpriseemployee.AlibabaeleenterpriseemployeebatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
