package newretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/newretail"
)

// AlibabaItApAddressSet setApAddressNew
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
func AlibabaItApAddressSet(clt *core.SDKClient, req *newretail.AlibabaItApAddressSetAPIRequest, session string) (*newretail.AlibabaItApAddressSetAPIResponse, error) {
	var resp newretail.AlibabaItApAddressSetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
