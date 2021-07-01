package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationServiceareaCheckAPIRequest
体检机构对接_上门检测服务范围查询 API请求
alibaba.alihealth.examination.servicearea.check

体检机构对接_上门检测服务范围查询 */
type AlibabaAlihealthExaminationServiceareaCheckAPIRequest struct {
	model.Params
	// 机构套餐编码
	_packageCode string
	// 上门检测地址
	_address string
	// 上门检测地址纬度
	_latitude string
	// 上门检测地址经度
	_longitude string
	// 省份名称（高德）
	_province string
	// 省份编码（高德adcode）
	_provinceCode string
	// 城市名称（高德）
	_city string
	// 城市编码（高德adcode）
	_cityCode string
	// 区域名称（高德）
	_district string
	// 区域编码（高德adcode）
	_districtCode string
}

// New
