package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationTodoorServiceinfoSync 上门检测服务信息同步
// alibaba.alihealth.examination.todoor.serviceinfo.sync
//
// isv同步上门检测服务信息给健康
func AlibabaAlihealthExaminationTodoorServiceinfoSync(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIRequest, resp *examination.AlibabaAlihealthExaminationTodoorServiceinfoSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
