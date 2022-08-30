package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusVisitorImageUpload 访客大厅图片上传及查看
// alibaba.campus.visitor.image.upload
//
// 访客大厅图片上传及查看
func AlibabaCampusVisitorImageUpload(clt *core.SDKClient, req *campus.AlibabaCampusVisitorImageUploadAPIRequest, session string) (*campus.AlibabaCampusVisitorImageUploadAPIResponse, error) {
	var resp campus.AlibabaCampusVisitorImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
