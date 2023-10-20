package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaofilesget 业务文件获取
// taobao.files.get
//
// 获取业务方暂存给ISV的文件列表
func Taobaofilesget(clt *core.SDKClient, req *util.TaobaofilesgetAPIRequest, session string) (*util.TaobaofilesgetAPIResponse, error) {
	var resp util.TaobaofilesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
