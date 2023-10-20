package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Taobaologisticsmediaresourcesupload 图片与视频上传
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
func Taobaologisticsmediaresourcesupload(clt *core.SDKClient, req *ascp.TaobaologisticsmediaresourcesuploadAPIRequest, session string) (*ascp.TaobaologisticsmediaresourcesuploadAPIResponse, error) {
	var resp ascp.TaobaologisticsmediaresourcesuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
