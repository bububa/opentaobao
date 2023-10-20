package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospayperiodagreementunsign 周期扣款协议解约接口
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
func Aliyunaliospayperiodagreementunsign(clt *core.SDKClient, req *aliospay.AliyunaliospayperiodagreementunsignAPIRequest, session string) (*aliospay.AliyunaliospayperiodagreementunsignAPIResponse, error) {
	var resp aliospay.AliyunaliospayperiodagreementunsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
