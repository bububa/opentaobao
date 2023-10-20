package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospayperiodagreementstatusget 查询周期扣款签约状态
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
func Aliyunaliospayperiodagreementstatusget(clt *core.SDKClient, req *aliospay.AliyunaliospayperiodagreementstatusgetAPIRequest, session string) (*aliospay.AliyunaliospayperiodagreementstatusgetAPIResponse, error) {
	var resp aliospay.AliyunaliospayperiodagreementstatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
