package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkPictureUpload 图片上传接口
// alibaba.wdk.picture.upload
//
// 上传图片
func AlibabaWdkPictureUpload(clt *core.SDKClient, req *wdkitem.AlibabaWdkPictureUploadAPIRequest, session string) (*wdkitem.AlibabaWdkPictureUploadAPIResponse, error) {
	var resp wdkitem.AlibabaWdkPictureUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
