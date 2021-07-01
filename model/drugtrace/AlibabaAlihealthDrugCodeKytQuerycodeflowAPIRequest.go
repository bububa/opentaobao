package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest
码流向查询 API请求
alibaba.alihealth.drug.code.kyt.querycodeflow

追溯码流向查询 */
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 追溯码
	_code string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 查询地所在省
	_queryProv string
	// 查询地所在市
	_queryCity string
	// 查询地所在区
	_queryArea string
	// 查询地所在区域代码
	_queryRegionCode string
	// 详细地址
	_detail string
}

// New
