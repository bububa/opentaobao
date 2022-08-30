package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJstSmsSignnameDelete 淘宝短信签名删除
// taobao.jst.sms.signname.delete
//
// 淘宝短信签名删除
func TaobaoJstSmsSignnameDelete(clt *core.SDKClient, req *jst.TaobaoJstSmsSignnameDeleteAPIRequest, session string) (*jst.TaobaoJstSmsSignnameDeleteAPIResponse, error) {
	var resp jst.TaobaoJstSmsSignnameDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
