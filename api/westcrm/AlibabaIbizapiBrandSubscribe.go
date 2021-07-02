package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaIbizapiBrandSubscribe 关注品牌号
// alibaba.ibizapi.brand.subscribe
//
// 关注品牌号服务
func AlibabaIbizapiBrandSubscribe(clt *core.SDKClient, req *westcrm.AlibabaIbizapiBrandSubscribeAPIRequest, session string) (*westcrm.AlibabaIbizapiBrandSubscribeAPIResponse, error) {
	var resp westcrm.AlibabaIbizapiBrandSubscribeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
