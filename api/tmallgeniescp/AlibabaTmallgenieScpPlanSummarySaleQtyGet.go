package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplansummarysaleqtyget 同步销售数据按照渠道类型汇总
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
func Alibabatmallgeniescpplansummarysaleqtyget(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplansummarysaleqtygetAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplansummarysaleqtygetAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplansummarysaleqtygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
