package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytsynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.synonymauths
//
// 物流企业查询货主企业信息
func Alibabaalihealthdrugkytsynonymauths(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytsynonymauthsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytsynonymauthsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytsynonymauthsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
