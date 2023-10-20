package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoPictureQnaigcUpload qnaigc业务线图片上传
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
func TaobaoPictureQnaigcUpload(clt *core.SDKClient, req *util.TaobaoPictureQnaigcUploadAPIRequest, resp *util.TaobaoPictureQnaigcUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
