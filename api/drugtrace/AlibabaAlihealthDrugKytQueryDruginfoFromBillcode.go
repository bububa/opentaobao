package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytquerydruginfofrombillcode 根据单据编号查询单据明细
// alibaba.alihealth.drug.kyt.query.druginfo.from.billcode
//
// 根据单据编号查询单据明细
func Alibabaalihealthdrugkytquerydruginfofrombillcode(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytquerydruginfofrombillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
