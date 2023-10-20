package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationserviceareacheck 体检机构对接_上门检测服务范围查询
// alibaba.alihealth.examination.servicearea.check
//
// 体检机构对接_上门检测服务范围查询
func Alibabaalihealthexaminationserviceareacheck(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationserviceareacheckAPIRequest, session string) (*examination.AlibabaalihealthexaminationserviceareacheckAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationserviceareacheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
