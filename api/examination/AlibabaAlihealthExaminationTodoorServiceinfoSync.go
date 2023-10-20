package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationtodoorserviceinfosync 上门检测服务信息同步
// alibaba.alihealth.examination.todoor.serviceinfo.sync
//
// isv同步上门检测服务信息给健康
func Alibabaalihealthexaminationtodoorserviceinfosync(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationtodoorserviceinfosyncAPIRequest, session string) (*examination.AlibabaalihealthexaminationtodoorserviceinfosyncAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationtodoorserviceinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
