package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmstemplatecreate 淘宝短信模板创建
// taobao.jst.sms.template.create
//
// 聚石塔短信模板创建
func Taobaojstsmstemplatecreate(clt *core.SDKClient, req *jst.TaobaojstsmstemplatecreateAPIRequest, session string) (*jst.TaobaojstsmstemplatecreateAPIResponse, error) {
	var resp jst.TaobaojstsmstemplatecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
