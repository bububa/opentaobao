package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeshopconfigdetailsubmit 店铺配置信息接口
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
func Alibabaalihousenewhomeshopconfigdetailsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeshopconfigdetailsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeshopconfigdetailsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeshopconfigdetailsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
