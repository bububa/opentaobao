package cainiaoncwl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoNcwlJhorderQueryAPIRequest
农村物流集货单查询接口 API请求
cainiao.ncwl.jhorder.query

提供给接入商家，查询农村物流集货单 */
type CainiaoNcwlJhorderQueryAPIRequest struct {
	model.Params
	// 1
	_param0 *JhRequest
}

// New
