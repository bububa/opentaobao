package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmactivitydataupdate 私域导购数据回流接口
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
func Alibabalsycrmactivitydataupdate(clt *core.SDKClient, req *tmallnr.AlibabalsycrmactivitydataupdateAPIRequest, session string) (*tmallnr.AlibabalsycrmactivitydataupdateAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmactivitydataupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
