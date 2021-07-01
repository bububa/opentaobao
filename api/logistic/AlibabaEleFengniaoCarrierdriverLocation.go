package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* AlibabaEleFengniaoCarrierdriverLocation
查询骑手当前位置
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置 */
func AlibabaEleFengniaoCarrierdriverLocation(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCarrierdriverLocationAPIRequest, session string) (*logistic.AlibabaEleFengniaoCarrierdriverLocationAPIResponse, error) {
	var resp logistic.AlibabaEleFengniaoCarrierdriverLocationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
