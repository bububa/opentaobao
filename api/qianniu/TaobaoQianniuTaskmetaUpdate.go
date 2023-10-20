package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskmetaupdate 更新任务元数据
// taobao.qianniu.taskmeta.update
//
// 由任务发起者调用
func Taobaoqianniutaskmetaupdate(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskmetaupdateAPIRequest, session string) (*qianniu.TaobaoqianniutaskmetaupdateAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskmetaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
