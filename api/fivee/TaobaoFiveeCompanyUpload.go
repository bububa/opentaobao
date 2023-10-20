package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeCompanyUpload 上传商信息接口
// taobao.fivee.company.upload
//
// 资质共享平台上传资质证照
func TaobaoFiveeCompanyUpload(clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyUploadAPIRequest, resp *fivee.TaobaoFiveeCompanyUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
