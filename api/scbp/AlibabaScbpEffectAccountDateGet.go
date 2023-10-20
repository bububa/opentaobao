package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpeffectaccountdateget 获取最近报表生成时间
// alibaba.scbp.effect.account.date.get
//
// 获取最近报表生成时间,格式为yyyy-MM-dd
func Alibabascbpeffectaccountdateget(clt *core.SDKClient, req *scbp.AlibabascbpeffectaccountdategetAPIRequest, session string) (*scbp.AlibabascbpeffectaccountdategetAPIResponse, error) {
	var resp scbp.AlibabascbpeffectaccountdategetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
