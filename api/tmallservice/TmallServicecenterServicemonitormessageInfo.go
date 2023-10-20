package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicemonitormessageInfo 服务预警单查询接口
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
func TmallServicecenterServicemonitormessageInfo(clt *core.SDKClient, req *tmallservice.TmallServicecenterServicemonitormessageInfoAPIRequest, session string) (*tmallservice.TmallServicecenterServicemonitormessageInfoAPIResponse, error) {
	var resp tmallservice.TmallServicecenterServicemonitormessageInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
