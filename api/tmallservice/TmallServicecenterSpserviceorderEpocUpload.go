package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderEpocUpload 电子保单文件上传接口
// tmall.servicecenter.spserviceorder.epoc.upload
//
// 电子保单文件上传接口
func TmallServicecenterSpserviceorderEpocUpload(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderEpocUploadAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderEpocUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
