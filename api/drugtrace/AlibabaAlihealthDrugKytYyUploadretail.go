package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytyyuploadretail 医院上传出库信息
// alibaba.alihealth.drug.kyt.yy.uploadretail
//
// 医院上传出库信息接口
func Alibabaalihealthdrugkytyyuploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytyyuploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytyyuploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytyyuploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
