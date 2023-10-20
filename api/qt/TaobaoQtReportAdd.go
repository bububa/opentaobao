package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaoqtreportadd 上传质检报告
// taobao.qt.report.add
//
// 上传质检报告
func Taobaoqtreportadd(clt *core.SDKClient, req *qt.TaobaoqtreportaddAPIRequest, session string) (*qt.TaobaoqtreportaddAPIResponse, error) {
	var resp qt.TaobaoqtreportaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
