package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsMediaResourcesUpload 图片与视频上传
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
func TaobaoLogisticsMediaResourcesUpload(clt *core.SDKClient, req *ascp.TaobaoLogisticsMediaResourcesUploadAPIRequest, session string) (*ascp.TaobaoLogisticsMediaResourcesUploadAPIResponse, error) {
	var resp ascp.TaobaoLogisticsMediaResourcesUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
