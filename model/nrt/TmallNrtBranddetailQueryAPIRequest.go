package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtBranddetailQueryAPIRequest
品牌详情查询 API请求
tmall.nrt.branddetail.query

根据品牌id查询品牌的详细信息 */
type TmallNrtBranddetailQueryAPIRequest struct {
	model.Params
	// 品牌id
	_brandId int64
}

// New
