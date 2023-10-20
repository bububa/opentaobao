package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrsearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.dr.searchstatus
//
// 单据处理状态查询
func Alibabaalihealthdrugkytdrsearchstatus(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrsearchstatusAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrsearchstatusAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrsearchstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
