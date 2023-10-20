package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaojstsmstemplatedelete 淘宝短信模板删除
// taobao.jst.sms.template.delete
//
// 淘宝短信模板删除
func Taobaojstsmstemplatedelete(clt *core.SDKClient, req *jst.TaobaojstsmstemplatedeleteAPIRequest, session string) (*jst.TaobaojstsmstemplatedeleteAPIResponse, error) {
	var resp jst.TaobaojstsmstemplatedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
