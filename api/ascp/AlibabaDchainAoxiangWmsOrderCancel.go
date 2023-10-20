package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwmsordercancel 回传发货单取消通知
// alibaba.dchain.aoxiang.wms.order.cancel
//
// 回传发货单取消通知
func Alibabadchainaoxiangwmsordercancel(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwmsordercancelAPIRequest, session string) (*ascp.AlibabadchainaoxiangwmsordercancelAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwmsordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
