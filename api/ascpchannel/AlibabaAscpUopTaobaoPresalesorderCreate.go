package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpuoptaobaopresalesordercreate 预售商家仓接单
// alibaba.ascp.uop.taobao.presalesorder.create
//
// 预售商家仓接单
func Alibabaascpuoptaobaopresalesordercreate(clt *core.SDKClient, req *ascpchannel.AlibabaascpuoptaobaopresalesordercreateAPIRequest, session string) (*ascpchannel.AlibabaascpuoptaobaopresalesordercreateAPIResponse, error) {
	var resp ascpchannel.AlibabaascpuoptaobaopresalesordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
