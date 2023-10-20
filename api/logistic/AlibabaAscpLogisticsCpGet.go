package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaascplogisticscpget 快递公司资源列表查询接口
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
func Alibabaascplogisticscpget(clt *core.SDKClient, req *logistic.AlibabaascplogisticscpgetAPIRequest, session string) (*logistic.AlibabaascplogisticscpgetAPIResponse, error) {
	var resp logistic.AlibabaascplogisticscpgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
