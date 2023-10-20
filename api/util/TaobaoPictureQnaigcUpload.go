package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaopictureqnaigcupload qnaigc业务线图片上传
// taobao.picture.qnaigc.upload
//
// qnaigc业务线图片上传
func Taobaopictureqnaigcupload(clt *core.SDKClient, req *util.TaobaopictureqnaigcuploadAPIRequest, session string) (*util.TaobaopictureqnaigcuploadAPIResponse, error) {
	var resp util.TaobaopictureqnaigcuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
