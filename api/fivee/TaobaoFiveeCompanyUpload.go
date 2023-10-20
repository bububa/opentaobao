package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveecompanyupload 上传商信息接口
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
func Taobaofiveecompanyupload(clt *core.SDKClient, req *fivee.TaobaofiveecompanyuploadAPIRequest, session string) (*fivee.TaobaofiveecompanyuploadAPIResponse, error) {
	var resp fivee.TaobaofiveecompanyuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
