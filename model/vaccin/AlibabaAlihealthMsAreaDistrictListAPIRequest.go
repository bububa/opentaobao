package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMsAreaDistrictListAPIRequest
疫苗预约地市信息查询 API请求
alibaba.alihealth.ms.area.district.list

微信小程序疫苗预约地市信息查询 */
type AlibabaAlihealthMsAreaDistrictListAPIRequest struct {
	model.Params
	// 省份ID
	_divisionId int64
}

// New
