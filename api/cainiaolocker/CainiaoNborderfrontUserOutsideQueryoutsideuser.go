package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaonborderfrontuseroutsidequeryoutsideuser 查询外部小件员休息
// cainiao.nborderfront.user.outside.queryoutsideuser
//
// 采用SPI方式查询外部公司的小件员信息
func Cainiaonborderfrontuseroutsidequeryoutsideuser(clt *core.SDKClient, req *cainiaolocker.CainiaonborderfrontuseroutsidequeryoutsideuserAPIRequest, session string) (*cainiaolocker.CainiaonborderfrontuseroutsidequeryoutsideuserAPIResponse, error) {
	var resp cainiaolocker.CainiaonborderfrontuseroutsidequeryoutsideuserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
