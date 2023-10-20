package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosfinancebankinfoquerybank 供应商银行账号查询
// alibaba.mos.finance.bankinfo.querybank
//
// 查询供应商对应的银行账号信息
func Alibabamosfinancebankinfoquerybank(clt *core.SDKClient, req *mos.AlibabamosfinancebankinfoquerybankAPIRequest, session string) (*mos.AlibabamosfinancebankinfoquerybankAPIResponse, error) {
	var resp mos.AlibabamosfinancebankinfoquerybankAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
