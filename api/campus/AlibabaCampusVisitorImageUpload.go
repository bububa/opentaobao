package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusvisitorimageupload 访客大厅图片上传及查看
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
func Alibabacampusvisitorimageupload(clt *core.SDKClient, req *campus.AlibabacampusvisitorimageuploadAPIRequest, session string) (*campus.AlibabacampusvisitorimageuploadAPIResponse, error) {
	var resp campus.AlibabacampusvisitorimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
