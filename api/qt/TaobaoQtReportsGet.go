package qt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qt"
)

// Taobaoqtreportsget 批量查询质检报告
// taobao.qt.reports.get
//
// 批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
func Taobaoqtreportsget(clt *core.SDKClient, req *qt.TaobaoqtreportsgetAPIRequest, session string) (*qt.TaobaoqtreportsgetAPIResponse, error) {
	var resp qt.TaobaoqtreportsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
