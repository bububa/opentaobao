package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljInstanceReport 淘宝客-推广者-淘礼金发放及使用报表
// taobao.tbk.dg.vegas.tlj.instance.report
//
// 淘礼金实例维度相关报表数据查询
func TaobaoTbkDgVegasTljInstanceReport(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljInstanceReportAPIRequest, session string) (*tbk.TaobaoTbkDgVegasTljInstanceReportAPIResponse, error) {
	var resp tbk.TaobaoTbkDgVegasTljInstanceReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
