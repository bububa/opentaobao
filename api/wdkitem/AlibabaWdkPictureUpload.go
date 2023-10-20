package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkPictureUpload 图片上传接口
// alibaba.wdk.picture.upload
//
// 上传图片
func AlibabaWdkPictureUpload(clt *core.SDKClient, req *wdkitem.AlibabaWdkPictureUploadAPIRequest, resp *wdkitem.AlibabaWdkPictureUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
