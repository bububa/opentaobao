package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytStorebilllistAPIRequest
零售端平台单据查询 API请求
alibaba.alihealth.drug.kyt.storebilllist

零售端平台单据查询 */
type AlibabaAlihealthDrugKytStorebilllistAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 单据状态(A全部 1上传成功 3处理成功 4处理失败 )
	_billStatus string
	// 页码
	_page int64
	// 页数
	_pageSize int64
}

// New
