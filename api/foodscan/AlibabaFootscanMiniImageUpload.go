package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// AlibabaFootscanMiniImageUpload 商家端图片上传
// alibaba.footscan.mini.image.upload
//
// 提供图片上传功能，同时进行图片的检测
func AlibabaFootscanMiniImageUpload(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniImageUploadAPIRequest, session string) (*foodscan.AlibabaFootscanMiniImageUploadAPIResponse, error) {
	var resp foodscan.AlibabaFootscanMiniImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
