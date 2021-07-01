package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintCustomaresGetAPIRequest
获取商家的自定义区模板信息 API请求
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息 */
type CainiaoCloudprintCustomaresGetAPIRequest struct {
	model.Params
	// 用户使用的标准模板id
	_templateId int64
}

// New
