package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationServiceareaCheck 体检机构对接_上门检测服务范围查询
// alibaba.alihealth.examination.servicearea.check
//
// 体检机构对接_上门检测服务范围查询
func AlibabaAlihealthExaminationServiceareaCheck(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationServiceareaCheckAPIRequest, resp *examination.AlibabaAlihealthExaminationServiceareaCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
