package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmsmessageshorturlcreate 聚石塔营销效果短链生成
// taobao.jst.sms.message.shorturl.create
//
// 聚石塔生成淘短链接口
func Taobaojstsmsmessageshorturlcreate(clt *core.SDKClient, req *jst.TaobaojstsmsmessageshorturlcreateAPIRequest, session string) (*jst.TaobaojstsmsmessageshorturlcreateAPIResponse, error) {
	var resp jst.TaobaojstsmsmessageshorturlcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
