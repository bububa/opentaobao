package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// Taobaoelifelifecardrecon 查询对账文件地址接口
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
func Taobaoelifelifecardrecon(clt *core.SDKClient, req *elife.TaobaoelifelifecardreconAPIRequest, session string) (*elife.TaobaoelifelifecardreconAPIResponse, error) {
	var resp elife.TaobaoelifelifecardreconAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
