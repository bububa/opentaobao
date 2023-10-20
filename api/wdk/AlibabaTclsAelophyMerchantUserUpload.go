package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantuserupload 商家会员数据上传
// alibaba.tcls.aelophy.merchant.user.upload
//
// 商家会员数据上传
func Alibabatclsaelophymerchantuserupload(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantuseruploadAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantuseruploadAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantuseruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
