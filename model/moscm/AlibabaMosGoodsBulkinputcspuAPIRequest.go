package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsBulkinputcspuAPIRequest
批量录入商品信息 API请求
alibaba.mos.goods.bulkinputcspu

用于商品信息的批量导入到银泰商品中台 */
type AlibabaMosGoodsBulkinputcspuAPIRequest struct {
	model.Params
	// 录入商品信息列表（最大列表长度：20）
	_cspuInputDtoList []CspuInputDto
}

// New
