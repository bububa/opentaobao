package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytwessynonymauths 物流企业查询货主企业信息
// alibaba.alihealth.drug.kyt.wes.synonymauths
//
// 物流企业查询货主企业信息
func Alibabaalihealthdrugkytwessynonymauths(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytwessynonymauthsAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytwessynonymauthsAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytwessynonymauthsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
