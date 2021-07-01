package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgContentGetAPIRequest
百川内容平台内容获取 API请求
alibaba.baichuan.ctg.content.get

百川内容平台内容获取 */
type AlibabaBaichuanCtgContentGetAPIRequest struct {
	model.Params
	// 投放位置
	_deliveryId string
	// 分页大小
	_pageSize int64
	// 当前页
	_currentPage int64
	// 资源位
	_resId string
	// 日期
	_date string
}

// New
