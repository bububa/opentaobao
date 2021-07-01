package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvGatewayAPIRequest
isv调用gateway API请求
alibaba.interact.isv.gateway

isv能够调用jae本身的server */
type AlibabaInteractIsvGatewayAPIRequest struct {
	model.Params
}

// New
