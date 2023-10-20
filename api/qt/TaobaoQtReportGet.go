package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaoqtreportget 查询质检报告
// taobao.qt.report.get
//
// 质检报告查询
func Taobaoqtreportget(clt *core.SDKClient, req *qt.TaobaoqtreportgetAPIRequest, session string) (*qt.TaobaoqtreportgetAPIResponse, error) {
	var resp qt.TaobaoqtreportgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
