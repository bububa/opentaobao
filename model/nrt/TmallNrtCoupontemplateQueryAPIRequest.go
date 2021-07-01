package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtCoupontemplateQueryAPIRequest
券模板查询 API请求
tmall.nrt.coupontemplate.query

新零售场景，商家拉取在新零售工作台设置的券数据 */
type TmallNrtCoupontemplateQueryAPIRequest struct {
	model.Params
	// 券列表
	_couponTypeList []int64
	// 当前页
	_currentPage int64
	// 每页数据数量
	_pageSize int64
	// 业务code阿里指定
	_bizCode string
}

// New
