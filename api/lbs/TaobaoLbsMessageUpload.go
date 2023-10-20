package lbs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lbs"
)

// Taobaolbsmessageupload lbs数据采集
// taobao.lbs.message.upload
//
// lbs数据采集
func Taobaolbsmessageupload(clt *core.SDKClient, req *lbs.TaobaolbsmessageuploadAPIRequest, session string) (*lbs.TaobaolbsmessageuploadAPIResponse, error) {
	var resp lbs.TaobaolbsmessageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
