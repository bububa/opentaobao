package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractWirelessDrawAPIRequest
双11到店互动无线端抽奖接口鉴权 API请求
alibaba.interact.wireless.draw

双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子 */
type AlibabaInteractWirelessDrawAPIRequest struct {
	model.Params
}

// NewAlibabaInteractWirelessDrawRequest 初始化AlibabaInteractWirelessDrawAPIRequest对象
func NewAlibabaInteractWirelessDrawRequest() *AlibabaInteractWirelessDrawAPIRequest {
	return &AlibabaInteractWirelessDrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractWirelessDrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.wireless.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractWirelessDrawAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
