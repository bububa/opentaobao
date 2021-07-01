package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariDisposeQueryAPIRequest
商品商家处置结果查询 API请求
taobao.sungari.dispose.query

红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理 */
type TaobaoSungariDisposeQueryAPIRequest struct {
	model.Params
	// 查询的key列表
	_paramList []string
}

// New
