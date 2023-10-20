package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangdeliverytemplatequery 商家运费模板查询
// alibaba.dchain.aoxiang.deliverytemplate.query
//
// 商家运费模板查询
func Alibabadchainaoxiangdeliverytemplatequery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangdeliverytemplatequeryAPIRequest, session string) (*ascp.AlibabadchainaoxiangdeliverytemplatequeryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangdeliverytemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
