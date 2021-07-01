package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkUmsRetrieveConfirm
回流单－外部对已拉取到的UMS单据进行确认
alibaba.wdk.ums.retrieve.confirm

回流单－外部对已拉取到的UMS单据进行确认 */
func AlibabaWdkUmsRetrieveConfirm(clt *core.SDKClient, req *wdk.AlibabaWdkUmsRetrieveConfirmAPIRequest, session string) (*wdk.AlibabaWdkUmsRetrieveConfirmAPIResponse, error) {
	var resp wdk.AlibabaWdkUmsRetrieveConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
