package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// Alibabanlifeb2citemdetailget b2c码详情查询
// alibaba.nlife.b2c.item.detail.get
//
// 根据零售+平台生成的唯一码获取对应详情
func Alibabanlifeb2citemdetailget(clt *core.SDKClient, req *nlife.Alibabanlifeb2citemdetailgetAPIRequest, session string) (*nlife.Alibabanlifeb2citemdetailgetAPIResponse, error) {
	var resp nlife.Alibabanlifeb2citemdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
