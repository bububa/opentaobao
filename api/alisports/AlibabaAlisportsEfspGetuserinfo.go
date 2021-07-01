package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

/* AlibabaAlisportsEfspGetuserinfo
获取用户详细信息
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息 */
func AlibabaAlisportsEfspGetuserinfo(clt *core.SDKClient, req *alisports.AlibabaAlisportsEfspGetuserinfoAPIRequest, session string) (*alisports.AlibabaAlisportsEfspGetuserinfoAPIResponse, error) {
	var resp alisports.AlibabaAlisportsEfspGetuserinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
