package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugUploadExtinfo
中药饮片及器械对接
alibaba.alihealth.drug.upload.extinfo

中药饮片及器械对接 */
func AlibabaAlihealthDrugUploadExtinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugUploadExtinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugUploadExtinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugUploadExtinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
