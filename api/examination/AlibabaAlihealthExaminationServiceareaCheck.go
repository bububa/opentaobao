package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

/* AlibabaAlihealthExaminationServiceareaCheck
体检机构对接_上门检测服务范围查询
alibaba.alihealth.examination.servicearea.check

体检机构对接_上门检测服务范围查询 */
func AlibabaAlihealthExaminationServiceareaCheck(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationServiceareaCheckAPIRequest, session string) (*examination.AlibabaAlihealthExaminationServiceareaCheckAPIResponse, error) {
	var resp examination.AlibabaAlihealthExaminationServiceareaCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
