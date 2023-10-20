package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderCheckstatusUpload 上传验收结果
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
func AlibabaIdleRentOrderCheckstatusUpload(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderCheckstatusUploadAPIRequest, resp *idle.AlibabaIdleRentOrderCheckstatusUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
