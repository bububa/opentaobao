package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallMsfIdentifyStatusQuery
喵师傅定案核销状态查询
tmall.msf.identify.status.query

喵师傅定案核销状态查询，供服务商erp系统调用 */
func TmallMsfIdentifyStatusQuery(clt *core.SDKClient, req *tmallservice.TmallMsfIdentifyStatusQueryAPIRequest, session string) (*tmallservice.TmallMsfIdentifyStatusQueryAPIResponse, error) {
	var resp tmallservice.TmallMsfIdentifyStatusQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
