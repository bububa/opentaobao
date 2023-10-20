package newretail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/newretail"
)

// Alibabaitapaddressget getApAddressByMacNew
// alibaba.it.ap.address.get
//
// 根据ap 的mac地址查询ap的结构化位置信息
func Alibabaitapaddressget(clt *core.SDKClient, req *newretail.AlibabaitapaddressgetAPIRequest, session string) (*newretail.AlibabaitapaddressgetAPIResponse, error) {
	var resp newretail.AlibabaitapaddressgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
