package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaalicomorderpreauthorizequeryfund 资金流水查询
// alibaba.alicom.order.preauthorize.query.fund
//
// 预授权-资金流水查询
func Alibabaalicomorderpreauthorizequeryfund(clt *core.SDKClient, req *alicom.AlibabaalicomorderpreauthorizequeryfundAPIRequest, session string) (*alicom.AlibabaalicomorderpreauthorizequeryfundAPIResponse, error) {
	var resp alicom.AlibabaalicomorderpreauthorizequeryfundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
