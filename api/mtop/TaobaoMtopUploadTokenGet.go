package mtop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtop"
)

// Taobaomtopuploadtokenget 获取文件上传授权
// taobao.mtop.upload.token.get
//
// 获取mtop文件上传授权
func Taobaomtopuploadtokenget(clt *core.SDKClient, req *mtop.TaobaomtopuploadtokengetAPIRequest, session string) (*mtop.TaobaomtopuploadtokengetAPIResponse, error) {
	var resp mtop.TaobaomtopuploadtokengetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
