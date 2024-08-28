package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStoreContractSyn 喵零合同同步
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
func TmallNrtStoreContractSyn(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtStoreContractSynAPIRequest, resp *nrt.TmallNrtStoreContractSynAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
