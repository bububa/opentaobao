package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuRfqdetailGetAPIRequest
获取RFQ详情 API请求
alibaba.icbu.rfqdetail.get

查看RFQ的详情信息 */
type AlibabaIcbuRfqdetailGetAPIRequest struct {
	model.Params
	// 验证
	_md5key string
	// 查询RFQ详情DTO
	_rfqQueryDto *RfqDetailSearchQueryDto
}

// New
