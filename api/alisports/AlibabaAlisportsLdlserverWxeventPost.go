package alisports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alisports"
)

/* AlibabaAlisportsLdlserverWxeventPost
乐动力线-微信设备绑定通知
alibaba.alisports.ldlserver.wxevent.post

转发弹外微信服务到弹内 */
func AlibabaAlisportsLdlserverWxeventPost(clt *core.SDKClient, req *alisports.AlibabaAlisportsLdlserverWxeventPostAPIRequest, session string) (*alisports.AlibabaAlisportsLdlserverWxeventPostAPIResponse, error) {
	var resp alisports.AlibabaAlisportsLdlserverWxeventPostAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
