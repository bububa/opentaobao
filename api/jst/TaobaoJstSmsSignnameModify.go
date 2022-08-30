package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameModify 淘宝短信签名修改
// taobao.jst.sms.signname.modify
//
// 淘宝短信签名修改，只能修改还未被审核的签名。
func TaobaoJstSmsSignnameModify(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameModifyAPIRequest, session string) (*jst.TaobaoJstSmsSignnameModifyAPIResponse, error) {
	var resp jst.TaobaoJstSmsSignnameModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
