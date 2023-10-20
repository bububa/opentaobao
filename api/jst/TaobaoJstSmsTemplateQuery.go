package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmstemplatequery 淘宝短信模板查询
// taobao.jst.sms.template.query
//
// 淘宝短信模板查询
func Taobaojstsmstemplatequery(clt *core.SDKClient, req *jst.TaobaojstsmstemplatequeryAPIRequest, session string) (*jst.TaobaojstsmstemplatequeryAPIResponse, error) {
	var resp jst.TaobaojstsmstemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
