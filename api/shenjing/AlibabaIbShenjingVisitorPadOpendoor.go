package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// Alibabaibshenjingvisitorpadopendoor 访客发起开门
// alibaba.ib.shenjing.visitor.pad.opendoor
//
// 访客PAD端录入完人脸后，可以点击开门按钮开门。
func Alibabaibshenjingvisitorpadopendoor(clt *core.SDKClient, req *shenjing.AlibabaibshenjingvisitorpadopendoorAPIRequest, session string) (*shenjing.AlibabaibshenjingvisitorpadopendoorAPIResponse, error) {
	var resp shenjing.AlibabaibshenjingvisitorpadopendoorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
