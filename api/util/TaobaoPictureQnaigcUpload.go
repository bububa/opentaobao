package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoPictureQnaigcUpload qnaigc业务线图片上传
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
func TaobaoPictureQnaigcUpload(clt *core.SDKClient, req *util.TaobaoPictureQnaigcUploadAPIRequest, session string) (*util.TaobaoPictureQnaigcUploadAPIResponse, error) {
	var resp util.TaobaoPictureQnaigcUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
