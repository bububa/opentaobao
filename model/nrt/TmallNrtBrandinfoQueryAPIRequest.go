package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtBrandinfoQueryAPIRequest
品牌数据查询 API请求
tmall.nrt.brandinfo.query

商家获取自己旗舰店授权的品牌id列表 */
type TmallNrtBrandinfoQueryAPIRequest struct {
	model.Params
}

// New
