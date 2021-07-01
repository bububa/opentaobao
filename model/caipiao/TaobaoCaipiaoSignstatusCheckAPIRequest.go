package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoSignstatusCheckAPIRequest
检查用户是否签署支付宝代购协议 API请求
taobao.caipiao.signstatus.check

检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。 */
type TaobaoCaipiaoSignstatusCheckAPIRequest struct {
	model.Params
}

// New
