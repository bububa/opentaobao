package ihome

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// TaobaoIhomeAdvancepicUpload ihome图片上传
// taobao.ihome.advancepic.upload
//
// ihome 定制业务编辑器投稿素材上传
func TaobaoIhomeAdvancepicUpload(clt *core.SDKClient, req *ihome.TaobaoIhomeAdvancepicUploadAPIRequest, session string) (*ihome.TaobaoIhomeAdvancepicUploadAPIResponse, error) {
	var resp ihome.TaobaoIhomeAdvancepicUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
