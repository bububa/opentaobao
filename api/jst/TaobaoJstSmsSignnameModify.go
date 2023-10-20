package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmssignnamemodify 淘宝短信签名修改
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
func Taobaojstsmssignnamemodify(clt *core.SDKClient, req *jst.TaobaojstsmssignnamemodifyAPIRequest, session string) (*jst.TaobaojstsmssignnamemodifyAPIResponse, error) {
	var resp jst.TaobaojstsmssignnamemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
