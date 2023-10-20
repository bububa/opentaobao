package legalcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// Alibabalegalcasecourttimeupdate 开庭时间变更
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
func Alibabalegalcasecourttimeupdate(clt *core.SDKClient, req *legalcase.AlibabalegalcasecourttimeupdateAPIRequest, session string) (*legalcase.AlibabalegalcasecourttimeupdateAPIResponse, error) {
	var resp legalcase.AlibabalegalcasecourttimeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
