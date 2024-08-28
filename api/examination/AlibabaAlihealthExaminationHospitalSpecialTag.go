package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthExaminationHospitalSpecialTag 体检机构获取特色服务标签
// alibaba.alihealth.examination.hospital.special.tag
//
// 体检机构获取特色服务标签列表
func AlibabaAlihealthExaminationHospitalSpecialTag(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationHospitalSpecialTagAPIRequest, resp *examination.AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
