package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodesellerwarehousesearch 查询仓库api
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
func Alibabaalihealthtracecodesellerwarehousesearch(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodesellerwarehousesearchAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodesellerwarehousesearchAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodesellerwarehousesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
