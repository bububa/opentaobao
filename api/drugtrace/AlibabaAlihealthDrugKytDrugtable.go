package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytdrugtable 查询药品目录信息
// alibaba.alihealth.drug.kyt.drugtable
//
// 查询药品目录信息
func Alibabaalihealthdrugkytdrugtable(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytdrugtableAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytdrugtableAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytdrugtableAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
