package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterspserviceorderepocupload 电子保单文件上传接口
// tmall.servicecenter.spserviceorder.epoc.upload
//
// 电子保单文件上传接口
func Tmallservicecenterspserviceorderepocupload(clt *core.SDKClient, req *tmallservice.TmallservicecenterspserviceorderepocuploadAPIRequest, session string) (*tmallservice.TmallservicecenterspserviceorderepocuploadAPIResponse, error) {
	var resp tmallservice.TmallservicecenterspserviceorderepocuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
