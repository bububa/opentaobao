package guoguo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/guoguo"
)

// CainiaoGuoguoCpBackupAssigncourierbyid 根据菜鸟账号ID指派小件员
// cainiao.guoguo.cp.backup.assigncourierbyid
//
// 根据菜鸟账号ID指派小件员
func CainiaoGuoguoCpBackupAssigncourierbyid(clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIRequest, session string) (*guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse, error) {
	var resp guoguo.CainiaoGuoguoCpBackupAssigncourierbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
