package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSdtcancel 通知速店通取消取号
// taobao.omniorder.store.sdtcancel
//
// 通知速店通取消取号
func TaobaoOmniorderStoreSdtcancel(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtcancelAPIRequest, session string) (*omniorder.TaobaoOmniorderStoreSdtcancelAPIResponse, error) {
	var resp omniorder.TaobaoOmniorderStoreSdtcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
