package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgvegastljreport 淘宝客-推广者-淘礼金效果数据
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
func Taobaotbkdgvegastljreport(clt *core.SDKClient, req *tbk.TaobaotbkdgvegastljreportAPIRequest, session string) (*tbk.TaobaotbkdgvegastljreportAPIResponse, error) {
	var resp tbk.TaobaotbkdgvegastljreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
