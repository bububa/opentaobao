package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaoqtreportupdate 更新质检报告
// taobao.qt.report.update
//
// 更新质检报告
func Taobaoqtreportupdate(clt *core.SDKClient, req *qt.TaobaoqtreportupdateAPIRequest, session string) (*qt.TaobaoqtreportupdateAPIResponse, error) {
	var resp qt.TaobaoqtreportupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
