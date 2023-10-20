package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// Alitriprailircarrierget 国际火车票铁路承运公司查询
// alitrip.rail.ir.carrier.get
//
// 国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
func Alitriprailircarrierget(clt *core.SDKClient, req *rail.AlitriprailircarriergetAPIRequest, session string) (*rail.AlitriprailircarriergetAPIResponse, error) {
	var resp rail.AlitriprailircarriergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
