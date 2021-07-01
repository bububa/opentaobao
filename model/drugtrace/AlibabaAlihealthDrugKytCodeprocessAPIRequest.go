package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytCodeprocessAPIRequest
关联关系处理查询 API请求
alibaba.alihealth.drug.kyt.codeprocess

关联关系处理查询 */
type AlibabaAlihealthDrugKytCodeprocessAPIRequest struct {
	model.Params
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
	// 上传标识
	_uploadFlag string
	// 处理状态
	_processFlag string
	// 批次号
	_produceBatchNo string
	// 查询标识
	_queryFlag string
	// 药品类型
	_physicType string
	// 生产企业ID
	_prodSeqNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 包装规格
	_pkgSpec string
	// 页数
	_page int64
	// 条数
	_pageSize int64
	// 客户端
	_clientType string
	// 企业ID
	_refEntId string
}

// New
