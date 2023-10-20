package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaoqtreportdelete 质检报告删除接口
// taobao.qt.report.delete
//
// 删除质检报告
func Taobaoqtreportdelete(clt *core.SDKClient, req *qt.TaobaoqtreportdeleteAPIRequest, session string) (*qt.TaobaoqtreportdeleteAPIResponse, error) {
	var resp qt.TaobaoqtreportdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
