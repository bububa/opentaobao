package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmstemplatemodify 淘宝短信模板修改
// taobao.jst.sms.template.modify
//
// 淘宝短信模板修改
func Taobaojstsmstemplatemodify(clt *core.SDKClient, req *jst.TaobaojstsmstemplatemodifyAPIRequest, session string) (*jst.TaobaojstsmstemplatemodifyAPIResponse, error) {
	var resp jst.TaobaojstsmstemplatemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
