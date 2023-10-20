package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// Alibabawdkpictureupload 图片上传接口
// alibaba.wdk.picture.upload
//
// 上传图片
func Alibabawdkpictureupload(clt *core.SDKClient, req *wdkitem.AlibabawdkpictureuploadAPIRequest, session string) (*wdkitem.AlibabawdkpictureuploadAPIResponse, error) {
	var resp wdkitem.AlibabawdkpictureuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
