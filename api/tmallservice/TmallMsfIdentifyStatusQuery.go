package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMsfIdentifyStatusQuery 喵师傅定案核销状态查询
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
func TmallMsfIdentifyStatusQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallMsfIdentifyStatusQueryAPIRequest, resp *tmallservice.TmallMsfIdentifyStatusQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
