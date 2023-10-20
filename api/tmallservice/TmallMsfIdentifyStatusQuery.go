package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMsfIdentifyStatusQuery 喵师傅定案核销状态查询
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
func TmallMsfIdentifyStatusQuery(clt *core.SDKClient, req *tmallservice.TmallMsfIdentifyStatusQueryAPIRequest, resp *tmallservice.TmallMsfIdentifyStatusQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
