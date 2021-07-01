package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugScanQuerycodeAPIRequest
查询药监码对应的有效期和包装规格 API请求
alibaba.alihealth.drug.scan.querycode

查询药监码对应的有效期和包装规格 */
type AlibabaAlihealthDrugScanQuerycodeAPIRequest struct {
	model.Params
	// 溯源码
	_code string
	// 用户标识id
	_webchatId string
	// 省编码
	_provinceCode string
	// 市编码
	_cityCode string
	// 区编码
	_areaCode string
	// 扫码日期
	_scanTime string
}

// New
