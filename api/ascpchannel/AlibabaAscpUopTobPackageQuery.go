package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuoptobpackagequery B2B包裹查询接口
// alibaba.ascp.uop.tob.package.query
//
// 供应链中台TOB包裹查询接口
func Alibabaascpuoptobpackagequery(clt *core.SDKClient, req *ascpchannel.AlibabaascpuoptobpackagequeryAPIRequest, session string) (*ascpchannel.AlibabaascpuoptobpackagequeryAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuoptobpackagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
