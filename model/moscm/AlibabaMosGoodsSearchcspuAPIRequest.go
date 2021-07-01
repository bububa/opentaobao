package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsSearchcspuAPIRequest
cspu查询 API请求
alibaba.mos.goods.searchcspu

商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流） */
type AlibabaMosGoodsSearchcspuAPIRequest struct {
	model.Params
	// 组合查询对象
	_paramCspuCriteria *CspuCriteria
	// 分页参数
	_paramPaginator *Paginator
}

// New
