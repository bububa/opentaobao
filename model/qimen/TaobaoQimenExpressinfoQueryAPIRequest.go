package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenExpressinfoQueryAPIRequest
配送公司信息查询接口 API请求
taobao.qimen.expressinfo.query

配送公司信息查询 */
type TaobaoQimenExpressinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenExpressinfoQueryRequest
}

// New
