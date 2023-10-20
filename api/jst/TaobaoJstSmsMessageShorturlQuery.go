package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmsmessageshorturlquery 聚石塔短链信息查询
// taobao.jst.sms.message.shorturl.query
//
// 聚石塔短链信息查询
func Taobaojstsmsmessageshorturlquery(clt *core.SDKClient, req *jst.TaobaojstsmsmessageshorturlqueryAPIRequest, session string) (*jst.TaobaojstsmsmessageshorturlqueryAPIResponse, error) {
	var resp jst.TaobaojstsmsmessageshorturlqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
