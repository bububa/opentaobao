package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoFilesGet 业务文件获取
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
func TaobaoFilesGet(clt *core.SDKClient, req *util.TaobaoFilesGetAPIRequest, resp *util.TaobaoFilesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
