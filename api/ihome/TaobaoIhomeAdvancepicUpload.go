package ihome

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// Taobaoihomeadvancepicupload ihome图片上传
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
func Taobaoihomeadvancepicupload(clt *core.SDKClient, req *ihome.TaobaoihomeadvancepicuploadAPIRequest, session string) (*ihome.TaobaoihomeadvancepicuploadAPIResponse, error) {
	var resp ihome.TaobaoihomeadvancepicuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
