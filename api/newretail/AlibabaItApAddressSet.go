package newretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/newretail"
)

// Alibabaitapaddressset setApAddressNew
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
func Alibabaitapaddressset(clt *core.SDKClient, req *newretail.AlibabaitapaddresssetAPIRequest, session string) (*newretail.AlibabaitapaddresssetAPIResponse, error) {
	var resp newretail.AlibabaitapaddresssetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
