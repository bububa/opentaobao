package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospayrecordlist 支付记录批量查询接口
// aliyun.alios.pay.record.list
//
// 商户用来对账的接口
func Aliyunaliospayrecordlist(clt *core.SDKClient, req *aliospay.AliyunaliospayrecordlistAPIRequest, session string) (*aliospay.AliyunaliospayrecordlistAPIResponse, error) {
	var resp aliospay.AliyunaliospayrecordlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
