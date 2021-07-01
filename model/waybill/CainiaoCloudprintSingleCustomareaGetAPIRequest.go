package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintSingleCustomareaGetAPIRequest
获取商家单一自定义区 API请求
cainiao.cloudprint.single.customarea.get

商家所有快递公司模板只有一个自定义区 */
type CainiaoCloudprintSingleCustomareaGetAPIRequest struct {
	model.Params
	// 这是商家用户id
	_sellerId int64
}

// New
