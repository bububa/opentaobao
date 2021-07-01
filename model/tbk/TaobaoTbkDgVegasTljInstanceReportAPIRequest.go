package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgVegasTljInstanceReportAPIRequest
淘宝客-推广者-淘礼金发放及使用报表 API请求
taobao.tbk.dg.vegas.tlj.instance.report

淘礼金实例维度相关报表数据查询 */
type TaobaoTbkDgVegasTljInstanceReportAPIRequest struct {
	model.Params
	// 实例ID
	_rightsId string
}

// New
