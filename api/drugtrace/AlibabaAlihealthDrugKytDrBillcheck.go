package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrbillcheck 疫苗追溯验证
// alibaba.alihealth.drug.kyt.dr.billcheck
//
// 各级疾控在入库完成后，需要做追溯信息验证
func Alibabaalihealthdrugkytdrbillcheck(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrbillcheckAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrbillcheckAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrbillcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
