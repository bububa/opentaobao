package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDruploadretail 快易通多融零售上传接口
// alibaba.alihealth.drug.kyt.druploadretail
//
// 快易通多融零售上传接口
func AlibabaAlihealthDrugKytDruploadretail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDruploadretailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDruploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDruploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
