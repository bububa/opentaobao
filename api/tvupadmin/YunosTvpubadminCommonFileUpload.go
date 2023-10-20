package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincommonfileupload 文件上传API
// yunos.tvpubadmin.common.file.upload
//
// 文件上传服务
func Yunostvpubadmincommonfileupload(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincommonfileuploadAPIRequest, session string) (*tvupadmin.YunostvpubadmincommonfileuploadAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincommonfileuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
