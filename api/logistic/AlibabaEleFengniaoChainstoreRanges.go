package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaochainstoreranges 蜂鸟查询门店配送范围接口
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
func Alibabaelefengniaochainstoreranges(clt *core.SDKClient, req *logistic.AlibabaelefengniaochainstorerangesAPIRequest, session string) (*logistic.AlibabaelefengniaochainstorerangesAPIResponse, error) {
	var resp logistic.AlibabaelefengniaochainstorerangesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
