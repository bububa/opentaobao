package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBeneiftDrawAPIRequest
抽奖接口 API请求
alibaba.beneift.draw

抽奖接口 */
type AlibabaBeneiftDrawAPIRequest struct {
	model.Params
	// 奖池唯一标识，奖池创建时即生成
	_ename string
	// 调用方AppName：规定为promotioncenter-${appId}
	_appName string
	// 调用方应用ip，非必填
	_ip string
}

// New
